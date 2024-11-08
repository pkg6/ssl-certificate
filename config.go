package certificate

import (
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

type Config struct {
	Domains      []string              `json:"domain" xml:"domain" yaml:"domain"`
	Registration *registrations.Config `json:"registration" xml:"registration" yaml:"registration"`
	Provider     *providers.Config     `json:"provider" xml:"provider" yaml:"provider"`
}

type DomainsDeploysConfig struct {
	Domains []*DomainDeployConfig `json:"domains" xml:"domains" yaml:"domains"`
	Deploys map[string]any        `json:"deploys" xml:"deploys" yaml:"deploys"`
}

type DomainDeployConfig struct {
	Deploy      string  `json:"deploy" xml:"deploy" yaml:"deploy"`
	Certificate *Config `json:"certificate" xml:"certificate" yaml:"certificate"`
}

func (d *DomainDeployConfig) DeployerConfig(deploys map[string]any) *deployer.Config {
	return deployer.MapNameAny(d.Deploy, deploys[d.Deploy])
}
