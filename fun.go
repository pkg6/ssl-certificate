package certificate

import (
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

type Config struct {
	Domains      []string              `json:"domain" xml:"domain" yaml:"domain"`
	Registration *registrations.Config `json:"registration" xml:"registration" yaml:"registration"`
	Provider     *providers.Config     `json:"provider" xml:"provider" yaml:"provider"`
}

func SSLCertificateByConfig(config *Config) (*registrations.Certificate, error) {
	provider, err := providers.NewProvider(config.Provider, config.Registration, config.Domains)
	if err != nil {
		return nil, err
	}
	return provider.Apply()
}

func SSLCertificate(email string, domain []string, provider string, providerConfig any) (*registrations.Certificate, error) {
	return SSLCertificateByConfig(&Config{
		Domains: domain,
		Provider: &providers.Config{
			Name:   provider,
			Config: providerConfig,
		},
		Registration: &registrations.Config{
			Email:    email,
			Provider: registrations.LetsencryptSSL,
		},
	})
}