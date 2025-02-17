package providers

import (
	namesiloProvider "github.com/go-acme/lego/v4/providers/dns/namesilo"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type NameSiloAccess struct {
	ApiKey string `json:"apiKey" xml:"apiKey" yaml:"apiKey"`
}

type NameSilo struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewNameSilo(options *Options) IProvider {
	return &NameSilo{Options: options}
}

func (n *NameSilo) Apply() (*registrations.Certificate, error) {
	access := &NameSiloAccess{}
	_ = pkg.JsonUnmarshal(n.Options.Config, access)
	_ = pkg.Setenv("NAMESILO_API_KEY", access.ApiKey)
	_ = pkg.SetTimeOut("NAMESILO_PROPAGATION_TIMEOUT")
	dnsProvider, err := namesiloProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(n.Options, dnsProvider)
}
