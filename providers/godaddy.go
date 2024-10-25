package providers

import (
	godaddyProvider "github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type GodaddyAccess struct {
	ApiKey    string `json:"apiKey" yaml:"apiKey" xml:"apiKey"`
	ApiSecret string `json:"apiSecret" yaml:"apiSecret" xml:"apiSecret"`
}
type godaddy struct {
	options *Options `json:"option" xml:"options" yaml:"options"`
}

func NewGodaddy(options *Options) IProvider {
	return &godaddy{options: options}
}

func (a *godaddy) Apply() (*registrations.Certificate, error) {

	access := &GodaddyAccess{}
	_ = helper.JsonUnmarshal(a.options.Config, access)

	_ = os.Setenv("GODADDY_API_KEY", access.ApiKey)
	_ = os.Setenv("GODADDY_API_SECRET", access.ApiSecret)

	dnsProvider, err := godaddyProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(a.options, dnsProvider)
}
