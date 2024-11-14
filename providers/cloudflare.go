package providers

import (
	cloudflareProvider "github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type CloudflareAccess struct {
	DnsApiToken string `json:"dnsApiToken" xml:"dnsApiToken" yaml:"dnsApiToken"`
}

type Cloudflare struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewCloudflare(options *Options) IProvider {
	return &Cloudflare{Options: options}
}

func (c *Cloudflare) Apply() (*registrations.Certificate, error) {
	access := &CloudflareAccess{}
	_ = pkg.JsonUnmarshal(c.Options.Config, access)
	_ = pkg.Setenv("CLOUDFLARE_DNS_API_TOKEN", access.DnsApiToken)
	_ = pkg.SetTimeOut("CLOUDFLARE_PROPAGATION_TIMEOUT")
	provider, err := cloudflareProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(c.Options, provider)
}
