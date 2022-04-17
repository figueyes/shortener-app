package config

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func GetConnection(username, password string) *kafka.Dialer {
	if len(username) == 0 || len(password) == 0 {
		return kafka.DefaultDialer
	}
	rootCAs, _ := x509.SystemCertPool()

	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	dialer := &kafka.Dialer{
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: username, // access key
			Password: password, // secret
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            rootCAs,
		},
	}
	return dialer
}
