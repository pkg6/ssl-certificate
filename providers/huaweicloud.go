package providers

import (
	huaweicloudProvider "github.com/go-acme/lego/v4/providers/dns/huaweicloud"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type HuaweiCloudAccess struct {
	Region          string `json:"region" yaml:"region" xml:"region"`
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId" xml:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey" xml:"secretAccessKey"`
}

type huaweicloud struct {
	option *Options
}

func NewHuaweiCloud(option *Options) IProvider {
	return &huaweicloud{
		option: option,
	}
}

func (t *huaweicloud) Apply() (*registrations.Certificate, error) {
	access := &HuaweiCloudAccess{}
	helper.JsonUnmarshal(t.option.Config, access)
	os.Setenv("HUAWEICLOUD_REGION", access.Region)
	os.Setenv("HUAWEICLOUD_ACCESS_KEY_ID", access.AccessKeyId)
	os.Setenv("HUAWEICLOUD_SECRET_ACCESS_KEY", access.SecretAccessKey)
	dnsProvider, err := huaweicloudProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(t.option, dnsProvider)
}
