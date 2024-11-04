package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
)

type AliYunAccess struct {
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" xml:"accessKeySecret" yaml:"accessKeySecret"`
}

func NewALiYun(option *Options) IProvider {
	return &ALiYun{
		Option: option,
	}
}

type ALiYun struct {
	Option *Options `json:"option" xml:"option" yaml:"option"`
}

func (a *ALiYun) Apply() (*registrations.Certificate, error) {
	access := &AliYunAccess{}
	_ = helper.JsonUnmarshal(a.Option.Config, access)
	_ = helper.Setenv("ALICLOUD_ACCESS_KEY", access.AccessKeyId)
	_ = helper.Setenv("ALICLOUD_SECRET_KEY", access.AccessKeySecret)
	_ = helper.SetTimeOut("ALICLOUD_PROPAGATION_TIMEOUT")
	dnsProvider, err := alidns.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(a.Option, dnsProvider)
}
