package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type PDNSAccess struct {
	ApiUrl string `json:"apiUrl" xml:"apiUrl" yaml:"apiUrl"`
	ApiKey string `json:"apiKey" xml:"apiKey" yaml:"apiKey"`
}

type Powerdns struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewPowerdns(options *Options) IProvider {
	return &Powerdns{Options: options}
}

func (a *Powerdns) Apply() (*registrations.Certificate, error) {
	access := &PDNSAccess{}
	_ = pkg.JsonUnmarshal(a.Options.Config, access)
	_ = pkg.Setenv("PDNS_API_URL", access.ApiUrl)
	_ = pkg.Setenv("PDNS_API_KEY", access.ApiKey)
	_ = pkg.SetTimeOut("PDNS_PROPAGATION_TIMEOUT")
	dnsProvider, err := pdns.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(a.Options, dnsProvider)
}
