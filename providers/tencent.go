package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type TencentAccess struct {
	SecretId  string `json:"secretId" yaml:"secretId" xml:"secretId"`
	SecretKey string `json:"secretKey" yaml:"secretKey" xml:"secretKey"`
}
type Tencent struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewTencent(options *Options) IProvider {
	return &Tencent{Options: options}
}

func (t *Tencent) Apply() (*registrations.Certificate, error) {
	access := &TencentAccess{}
	_ = pkg.JsonUnmarshal(t.Options.Config, access)
	_ = pkg.Setenv("TENCENTCLOUD_SECRET_ID", access.SecretId)
	_ = pkg.Setenv("TENCENTCLOUD_SECRET_KEY", access.SecretKey)
	_ = pkg.SetTimeOut("TENCENTCLOUD_PROPAGATION_TIMEOUT")
	dnsProvider, err := tencentcloud.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(t.Options, dnsProvider)
}
