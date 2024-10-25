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
	options *Options `json:"option" xml:"option" yaml:"option"`
}

func NewCloudflare(options *Options) IProvider {
	return &cloudflare{options: options}
}

func (c *cloudflare) Apply() (*registrations.Certificate, error) {
	access := &CloudflareAccess{}
	_ = helper.JsonUnmarshal(c.options.Config, access)
	_ = os.Setenv("CLOUDFLARE_DNS_API_TOKEN", access.DnsApiToken)
	provider, err := cloudflareProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(c.options, provider)
}
