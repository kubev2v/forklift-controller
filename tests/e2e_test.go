package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-logr/logr"
	apis "github.com/konveyor/forklift-controller/pkg/apis"
	"github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	"github.com/konveyor/forklift-controller/pkg/settings"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func TestOvirtProvider(t *testing.T) {
	var Settings = &settings.Settings
	var log logr.Logger

	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println("unable to set up client config")
		os.Exit(1)
	}

	logf.SetLogger(
		logf.ZapLogger(Settings.Logging.Development))
	log = logf.Log.WithName("entrypoint")

	username := os.Getenv("OVIRT_USERNAME")
	if username == "" {
		t.Fatal("OVIRT_USERNAME is not set")
	}
	password := os.Getenv("OVIRT_PASSWORD")
	if password == "" {
		t.Fatal("OVIRT_PASSWORD is not set")
	}

	ovirtURL := os.Getenv("OVIRT_URL")
	if ovirtURL == "" {
		t.Fatal("OVIRT_URL is not set")
	}

	cacertFile := os.Getenv("OVIRT_CACERT")
	if cacertFile == "" {
		t.Fatal("OVIRT_CACERT is not set")
	}

	cacert, err := os.ReadFile(cacertFile)
	if err != nil {
		t.Fatalf("Could not read %s", cacertFile)
	}

	// Register
	v1beta1.SchemeBuilder.AddToScheme(scheme.Scheme)
	apis.AddToScheme(scheme.Scheme)

	cl, err := client.New(conf, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		t.Error(err, "Failed to create client")
	}

	log.Info("Creating namespace...")
	namespaceName := fmt.Sprintf("provider-test-ns-%d", time.Now().Unix())

	namespace := &corev1.Namespace{
		TypeMeta: v1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceName,
		},
	}

	defer cl.Delete(context.TODO(), namespace, &client.DeleteOptions{})

	err = cl.Create(context.TODO(), namespace, &client.CreateOptions{})
	if err != nil {
		t.Fatal(err, "Failed to create namespace")
	}

	log.Info("Creating secret...")

	secret := &corev1.Secret{
		TypeMeta: v1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespaceName,
			Name:      "ovirt-provider-test-secret",
		},
		Data: map[string][]byte{
			"user":     []byte(username),
			"password": []byte(password),
			"cacert":   cacert,
		},
		Type: corev1.SecretTypeOpaque,
	}
	err = cl.Create(context.TODO(), secret, &client.CreateOptions{})
	if err != nil {
		t.Fatal(err, "Failed to create secret")
	}

	providerName := v1.ObjectMeta{
		Namespace: namespaceName,
		Name:      "ovirt-vm",
	}

	ovirtProvider := v1beta1.OVirt
	p := &v1beta1.Provider{
		TypeMeta: v1.TypeMeta{
			Kind:       "Provider",
			APIVersion: "forklift.konveyor.io/v1beta1",
		},
		ObjectMeta: providerName,
		Spec: v1beta1.ProviderSpec{
			Type: &ovirtProvider,
			URL:  os.Getenv("OVIRT_URL"),
			Secret: corev1.ObjectReference{
				Name:      secret.Name,
				Namespace: namespaceName,
			},
		},
	}

	err = cl.Create(context.TODO(), p, &client.CreateOptions{})
	if err != nil {
		t.Fatal(err, "Failed to create provider")
	}

	returnedProvider := &v1beta1.Provider{}
	providerIdentifier := types.NamespacedName{Namespace: providerName.Namespace, Name: providerName.Name}
	err = cl.Get(context.TODO(), providerIdentifier, returnedProvider)
	if err != nil {
		t.Fatal(err, "Could not get provider")
	}

	done := make(chan struct{})

	// Wait for provider to be ready
	statusCheck := func() {
		err = cl.Get(context.TODO(), providerIdentifier, returnedProvider)
		if err != nil {
			t.Error(err, "Could not get provider")
		}

		if returnedProvider.Status.Conditions.IsReady() {
			log.Info("Provider is ready")
			close(done)
		}
	}

	go wait.Until(statusCheck, time.Second, done)

	timeout := time.After(1 * time.Minute)
	select {
	case <-timeout:
		t.Errorf("Provider is not ready in time, last status: %v", returnedProvider.Status)
	case <-done:
	}
}
