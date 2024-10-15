package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type AliYunAccess struct {
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" xml:"accessKeySecret" yaml:"accessKeySecret"`
}

func NewAliyun(option *Options) IProvider {
	return &aliyunApply{
		option: option,
	}
}

type aliyunApply struct {
	option *Options
}

func (a *aliyunApply) Apply() (*registrations.Certificate, error) {
	access := &AliYunAccess{}
	helper.JsonUnmarshal(a.option.Config, access)
	_ = os.Setenv("ALICLOUD_ACCESS_KEY", access.AccessKeyId)
	_ = os.Setenv("ALICLOUD_SECRET_KEY", access.AccessKeySecret)
	dnsProvider, err := alidns.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(a.option, dnsProvider)
}
