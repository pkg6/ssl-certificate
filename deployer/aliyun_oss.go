package deployer

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
)

type ALiYunOSSAccess struct {
	Endpoint        string `json:"endpoint" xml:"endpoint" yaml:"endpoint"`
	Bucket          string `json:"bucket" xml:"bucket" yaml:"bucket"`
	Domain          string `json:"domain" xml:"domain" yaml:"domain"`
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" xml:"accessKeySecret" yaml:"accessKeySecret"`
}

type aLiYunOSS struct {
	options *Options
	logs    []string

	access *ALiYunOSSAccess
	client *oss.Client
}

func NewALiYunOSS(options *Options) (IDeployer, error) {
	access := &ALiYunOSSAccess{}
	_ = helper.JsonUnmarshal(options.Access, access)
	a := &aLiYunOSS{
		options: options,
		access:  access,
		logs:    make([]string, 0),
	}
	client, err := oss.New(access.Endpoint, access.AccessKeyId, access.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	a.client = client
	return a, nil
}

func (d *aLiYunOSS) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	domain := d.access.Domain
	if domain == "" {
		domain = certificate.Domain
	}
	err := d.
		client.
		PutBucketCnameWithCertificate(d.access.Bucket, oss.PutBucketCname{
			Cname: domain,
			CertificateConfiguration: &oss.CertificateConfiguration{
				Certificate: certificate.Certificate,
				PrivateKey:  certificate.PrivateKey,
				Force:       true,
			},
		})
	if err != nil {
		return fmt.Errorf("deploy aliyun oss error: %w", err)
	}
	d.logs = append(d.logs, AddLog(OSS, "Deployment successful", nil))
	return nil
}

func (d *aLiYunOSS) GetLogs() []string {
	return d.logs
}
