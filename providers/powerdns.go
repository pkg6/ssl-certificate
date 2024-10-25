package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type PDNSAccess struct {
	ApiUrl string `json:"apiUrl" xml:"apiUrl" yaml:"apiUrl"`
	ApiKey string `json:"apiKey" xml:"apiKey" yaml:"apiKey"`
}

type powerdns struct {
	options *Options
}

func NewPowerdns(options *Options) IProvider {
	return &powerdns{options: options}
}

func (a *powerdns) Apply() (*registrations.Certificate, error) {
	access := &PDNSAccess{}
	_ = helper.JsonUnmarshal(a.options.Config, access)
	os.Setenv("PDNS_API_URL", access.ApiUrl)
	os.Setenv("PDNS_API_KEY", access.ApiKey)
	dnsProvider, err := pdns.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(a.options, dnsProvider)
}
