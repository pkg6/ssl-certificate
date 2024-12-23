package providers

import (
	huaweicloudProvider "github.com/go-acme/lego/v4/providers/dns/huaweicloud"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type HuaweiCloudAccess struct {
	Region          string `json:"region" yaml:"region" xml:"region"`
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId" xml:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey" xml:"secretAccessKey"`
}

type HuaWei struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewHuaweiCloud(options *Options) IProvider {
	return &HuaWei{Options: options}
}

func (t *HuaWei) Apply() (*registrations.Certificate, error) {
	access := &HuaweiCloudAccess{}
	_ = pkg.JsonUnmarshal(t.Options.Config, access)
	_ = pkg.Setenv("HUAWEICLOUD_REGION", access.Region)
	_ = pkg.Setenv("HUAWEICLOUD_ACCESS_KEY_ID", access.AccessKeyId)
	_ = pkg.Setenv("HUAWEICLOUD_SECRET_ACCESS_KEY", access.SecretAccessKey)
	_ = pkg.SetTimeOut("HUAWEICLOUD_PROPAGATION_TIMEOUT")
	dnsProvider, err := huaweicloudProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(t.Options, dnsProvider)
}
