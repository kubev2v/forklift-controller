package settings

import (
	"os"
	"strconv"
)

//
// Environment variables.
const (
	AdmissionWebhookEnabled = "ADMISSION_WEBHOOK_ENABLED"
	AdmissionWebhookPort    = "ADMISSION_WEBHOOK_PORT"
	AdmissionWebhookCertDir = "ADMISSION_WEBHOOK_CERTDIR"
)

//
// Admission webhook settings.
type AdmissionWebhook struct {
	// URL.
	Port int
	// Dir with tls.key and tls.crt
	CertDir string
	// Enabled
	Enabled bool
}

//
// Load settings.
func (r *AdmissionWebhook) Load() (err error) {
	// Admission webhook controller port
	r.Enabled = getEnvBool(AdmissionWebhookEnabled, false)
	if r.Enabled {
		if s, found := os.LookupEnv(AdmissionWebhookPort); found {
			r.Port, _ = strconv.Atoi(s)
		} else {
			r.Port = 8444
		}
		if s, found := os.LookupEnv(AdmissionWebhookCertDir); found {
			r.CertDir = s
		} else {
			r.CertDir = "/var/run/secrets/forklift-admission-webhook-serving-cert/"
		}
	}
	return
}
