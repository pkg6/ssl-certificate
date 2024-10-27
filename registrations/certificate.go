package registrations

import "github.com/go-acme/lego/v4/certificate"

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
