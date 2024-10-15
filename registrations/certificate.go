package registrations

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
