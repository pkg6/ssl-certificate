package providers

import (
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
)

type AwsAccess struct {
	Region          string `json:"region" xml:"region" yaml:"region"`
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" xml:"secretAccessKey" yaml:"secretAccessKey"`
	HostedZoneId    string `json:"hostedZoneId" xml:"hostedZoneId" yaml:"hostedZoneId"`
}

type AWS struct {
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

func NewAws(options *Options) IProvider {
	return &AWS{Options: options}
}

func (t *AWS) Apply() (*registrations.Certificate, error) {
	access := &AwsAccess{}
	_ = pkg.JsonUnmarshal(t.Options.Config, access)
	_ = pkg.Setenv("AWS_REGION", access.Region)
	_ = pkg.Setenv("AWS_ACCESS_KEY_ID", access.AccessKeyId)
	_ = pkg.Setenv("AWS_SECRET_ACCESS_KEY", access.SecretAccessKey)
	_ = pkg.Setenv("AWS_HOSTED_ZONE_ID", access.HostedZoneId)
	_ = pkg.SetTimeOut("AWS_PROPAGATION_TIMEOUT")
	dnsProvider, err := route53.NewDNSProvider()
	if err != nil {
		return nil, err
	}
	return Apply(t.Options, dnsProvider)
}
