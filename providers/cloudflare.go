package providers

import (
	cloudflareProvider "github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type CloudflareAccess struct {
	DnsApiToken string `json:"dnsApiToken" xml:"dnsApiToken" yaml:"dnsApiToken"`
}

type cloudflare struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewCloudflare(options *Options) IProvider {
	return &cloudflare{Options: options}
}

func (c *cloudflare) Apply() (*registrations.Certificate, error) {
	access := &CloudflareAccess{}
	_ = helper.JsonUnmarshal(c.Options.Config, access)
	_ = os.Setenv("CLOUDFLARE_DNS_API_TOKEN", access.DnsApiToken)
	provider, err := cloudflareProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(c.Options, provider)
}
