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
	option *Options
}

func NewGodaddy(option *Options) IProvider {
	return &godaddy{
		option: option,
	}
}

func (a *godaddy) Apply() (*registrations.Certificate, error) {

	access := &GodaddyAccess{}
	helper.JsonUnmarshal(a.option.Config, access)

	os.Setenv("GODADDY_API_KEY", access.ApiKey)
	os.Setenv("GODADDY_API_SECRET", access.ApiSecret)

	dnsProvider, err := godaddyProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}

	return apply(a.option, dnsProvider)
}
