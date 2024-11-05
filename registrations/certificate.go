package registrations

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-acme/lego/v4/certificate"
)

type Certificate struct {
	User              *User  `json:"user" xml:"user"`
	Domain            string `json:"domain" xml:"domain"`
	CertUrl           string `json:"certUrl" xml:"certUrl"`
	CertStableUrl     string `json:"certStableUrl" xml:"certStableUrl"`
	PrivateKey        string `json:"privateKey" xml:"privateKey"`
	Certificate       string `json:"certificate" xml:"certificate"`
	IssuerCertificate string `json:"issuerCertificate" xml:"issuerCertificate"`
	Csr               string `json:"csr" xml:"csr"`
}

func NewCertificateByResource(user *User, cert *certificate.Resource) *Certificate {
	return &Certificate{
		User:              user,
		Domain:            cert.Domain,
		CertUrl:           cert.CertURL,
		CertStableUrl:     cert.CertStableURL,
		PrivateKey:        string(cert.PrivateKey),
		Certificate:       string(cert.Certificate),
		IssuerCertificate: string(cert.IssuerCertificate),
		Csr:               string(cert.CSR),
	}
}

func (c *Certificate) X509Certificate() (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(c.Certificate))
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

func (c *Certificate) X509PrivateKey() (any, error) {
	block, _ := pem.Decode([]byte(c.PrivateKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse private key PEM")
	}
	// Check if it is an RSA private key
	if block.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	// Check if it is an ECDSA private key
	if block.Type == "EC PRIVATE KEY" {
		return x509.ParseECPrivateKey(block.Bytes)
	}
	return nil, fmt.Errorf("unsupported private key type")
}
