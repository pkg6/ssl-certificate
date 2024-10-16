package certificate

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
	"net/url"
	"strings"
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

// DomainCertificates
// Obtain domain certificate information
func DomainCertificates(domain string) ([]*CertificateInfo, error) {
	if strings.HasPrefix(domain, "https://") || strings.HasPrefix(domain, "http://") {
		parse, err := url.Parse(domain)
		if err != nil {
			return nil, err
		}
		domain = parse.Host
	}
	dial, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), nil)
	if err != nil {
		return nil, err
	}
	state := dial.ConnectionState()
	var infos []*CertificateInfo
	const blockType = "PUBLIC KEY"
	for _, certificate := range state.PeerCertificates {
		c := &CertificateInfo{
			Subject:            certificate.Subject.String(),
			Issuer:             certificate.Issuer.String(),
			NotBefore:          certificate.NotBefore,
			NotAfter:           certificate.NotAfter,
			PublicKeyAlgorithm: certificate.PublicKeyAlgorithm,
			Version:            certificate.Version,
		}
		pubKey, _ := x509.MarshalPKIXPublicKey(certificate.PublicKey)
		pemBlock := &pem.Block{Type: blockType, Bytes: pubKey}
		c.PublicKey = string(pem.EncodeToMemory(pemBlock))
		infos = append(infos, c)
	}
	return infos, nil
}
