package certificate

import (
	"context"
	"github.com/go-acme/lego/v4/log"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

// SSLCertificateByConfig
// Generate certificate
func SSLCertificateByConfig(config *Config) (*registrations.Certificate, error) {
	provider, err := providers.NewProvider(config.Provider, config.Registration, config.Domains)
	if err != nil {
		return nil, err
	}
	return provider.Apply()
}

// SSLCertificate
// Generate certificate
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

// Deployer
// Deploy the certificate
func Deployer(ctx context.Context, config *deployer.Config, certificate *registrations.Certificate) ([]string, error) {
	return deployer.Run(ctx, config, certificate)
}

// SSLCertificateDeployer
// Certificate generation and automatic deployment completion
func SSLCertificateDeployer(ctx context.Context, cfg *Config, deployer *deployer.Config) error {
	certificate, err := SSLCertificateByConfig(cfg)
	if err != nil {
		log.Fatalf("Generate SSL Certificate err=`%v`", err)
		return err
	}
	logs, err := Deployer(ctx, deployer, certificate)
	if err != nil {
		log.Fatalf("Deploy err=%v", err)
		return err
	}
	for _, l := range logs {
		log.Println(l)
	}
	return nil
}
