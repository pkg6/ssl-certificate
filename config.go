package certificate

import (
	"crypto/x509"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
	"time"
)

type Config struct {
	Domains      []string              `json:"domain" xml:"domain" yaml:"domain"`
	Registration *registrations.Config `json:"registration" xml:"registration" yaml:"registration"`
	Provider     *providers.Config     `json:"provider" xml:"provider" yaml:"provider"`
}

type DomainsDeploysConfig struct {
	Domains []*DomainDeployConfig `json:"domains"`
	Deploys map[string]any        `json:"deploys"`
}

type DomainDeployConfig struct {
	Deploy      string  `json:"deploy"`
	Certificate *Config `json:"certificate"`
}

func (d *DomainDeployConfig) DeployerConfig(deploys map[string]any) *deployer.Config {
	return deployer.MapNameAny(d.Deploy, deploys[d.Deploy])
}

type CertificateInfo struct {
	Subject            string                  `json:"subject" xml:"subject"`
	Issuer             string                  `json:"issuer" xml:"issuer"`
	NotBefore          time.Time               `json:"not_before" xml:"NotBefore"`
	NotAfter           time.Time               `json:"not_after" xml:"NotAfter"`
	PublicKeyAlgorithm x509.PublicKeyAlgorithm `json:"public_key_algorithm" xml:"PublicKeyAlgorithm"`
	Version            int                     `json:"version" xml:"version"`
	PublicKey          string                  `json:"public_key" xml:"PublicKey"`
}
