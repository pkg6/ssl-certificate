package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
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
	_ = helper.JsonUnmarshal(t.Options.Config, access)
	_ = os.Setenv("TENCENTCLOUD_SECRET_ID", access.SecretId)
	_ = os.Setenv("TENCENTCLOUD_SECRET_KEY", access.SecretKey)
	dnsProvider, err := tencentcloud.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(t.Options, dnsProvider)
}
