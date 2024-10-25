package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
)

type AwsAccess struct {
	Region          string `json:"region" xml:"region" yaml:"region"`
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" xml:"secretAccessKey" yaml:"secretAccessKey"`
	HostedZoneId    string `json:"hostedZoneId" xml:"hostedZoneId" yaml:"hostedZoneId"`
}

type aws struct {
	options *Options
}

func NewAws(options *Options) IProvider {
	return &aws{
		options: options,
	}
}

func (t *aws) Apply() (*registrations.Certificate, error) {
	access := &AwsAccess{}
	_ = helper.JsonUnmarshal(t.options.Config, access)
	os.Setenv("AWS_REGION", access.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", access.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", access.SecretAccessKey)
	os.Setenv("AWS_HOSTED_ZONE_ID", access.HostedZoneId)
	dnsProvider, err := route53.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return apply(t.options, dnsProvider)
}
