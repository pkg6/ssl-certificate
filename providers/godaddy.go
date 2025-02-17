package providers

import (
	godaddyProvider "github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type GodaddyAccess struct {
	ApiKey    string `json:"apiKey" yaml:"apiKey" xml:"apiKey"`
	ApiSecret string `json:"apiSecret" yaml:"apiSecret" xml:"apiSecret"`
}
type Godaddy struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewGodaddy(options *Options) IProvider {
	return &Godaddy{Options: options}
}

func (a *Godaddy) Apply() (*registrations.Certificate, error) {

	access := &GodaddyAccess{}
	_ = pkg.JsonUnmarshal(a.Options.Config, access)

	_ = pkg.Setenv("GODADDY_API_KEY", access.ApiKey)
	_ = pkg.Setenv("GODADDY_API_SECRET", access.ApiSecret)
	_ = pkg.SetTimeOut("GODADDY_PROPAGATION_TIMEOUT")
	dnsProvider, err := godaddyProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(a.Options, dnsProvider)
}
