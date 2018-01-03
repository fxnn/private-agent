package config

import "time"

const (
	DefaultMaxWorkerTimeoutInSeconds  = 100
	DefaultMaxRequestTimeoutInSeconds = 24 * 60 * 60
	DefaultCertificateValidFor        = 365 * 24 * time.Hour
)

var DefaultCertificateHosts = []string{"localhost", "127.0.0.1"}

// Drop configuration, created once per configured drop
type Drop struct {
	// Name identifies this drop uniquely
	Name string

	// ListenAddress defines the the local network address this drop
	// shall listen on.
	ListenAddress string

	// MaxWorkerTimeoutInSeconds limits the time period a worker registration may be considered active without having
	// received an update from the worker.
	MaxWorkerTimeoutInSeconds int

	// MaxRequestTimeoutInSeconds limits the time period during which a request and its response may be processed and
	// retrieved.
	MaxRequestTimeoutInSeconds int

	// PrivateKeySize is the size of the private RSA key in bits, mostly 2048 oder 4096.
	PrivateKeySize int

	// CertificateValidFor determines how long a autogenerated certificate is valid, starting from now.
	CertificateValidFor time.Duration

	// CertificateHosts is a list of IP addresses or hostnames a autogenerated certificate is valid for.
	CertificateHosts []string
}
