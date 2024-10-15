package certificate

import (
	"context"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

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

func Deployer(config *deployer.Config, ctx context.Context, certificate *registrations.Certificate) ([]string, error) {
	dep, err := deployer.NewDeployer(config)
	if err != nil {
		return nil, err
	}
	if err := dep.Deploy(ctx, certificate); err != nil {
		return nil, err
	}
	return dep.GetLogs(), err
}
