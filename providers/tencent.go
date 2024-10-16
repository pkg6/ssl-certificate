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
type tencent struct {
	option *Options
}

func NewTencent(option *Options) IProvider {
	return &tencent{
		option: option,
	}
}

func (t *tencent) Apply() (*registrations.Certificate, error) {
	access := &TencentAccess{}
	_ = helper.JsonUnmarshal(t.option.Config, access)
	_ = os.Setenv("TENCENTCLOUD_SECRET_ID", access.SecretId)
	_ = os.Setenv("TENCENTCLOUD_SECRET_KEY", access.SecretKey)
	dnsProvider, err := tencentcloud.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(t.option, dnsProvider)
}
