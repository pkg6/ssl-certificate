package providers

import (
	cloudflareProvider "github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type CloudflareAccess struct {
	DnsApiToken string `json:"dnsApiToken" yaml:"dnsApiToken" xml:"dnsApiToken"`
}

type cloudflare struct {
	option *Options
}

func NewCloudflare(option *Options) IProvider {
	return &cloudflare{option: option}
}

func (c *cloudflare) Apply() (*registrations.Certificate, error) {
	access := &CloudflareAccess{}
	helper.JsonUnmarshal(c.option.Config, access)
	os.Setenv("CLOUDFLARE_DNS_API_TOKEN", access.DnsApiToken)
	provider, err := cloudflareProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(c.option, provider)
}
