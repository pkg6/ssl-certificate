package deployer

import (
	"context"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg6/ssl-certificate/registrations"
)

type ALiYunCDNAccess struct {
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" xml:"accessKeySecret" yaml:"accessKeySecret"`
	Region          string `json:"region" xml:"region" yaml:"region"`
	Endpoint        string `json:"endpoint" xml:"endpoint" yaml:"endpoint"`
	Domain          string `json:"domain" xml:"domain" yaml:"domain"`
}

type aLiYunCDN struct {
	cdn20180510 *cdn20180510.Client
	options     *Options
	access      *ALiYunCDNAccess
	logs        []string
}

func NewALiYunCDN(options *Options) (IDeployer, error) {
	access := &ALiYunCDNAccess{}
	_ = options.JsonUnmarshal(access)
	if access.Region == "" {
		access.Region = "cn-hangzhou"
	}
	if access.Endpoint == "" {
		access.Endpoint = "cdn.aliyuncs.com"
	}
	a := &aLiYunCDN{
		options: options,
		access:  access,
	}
	client, err := a.createClient(access.AccessKeyId, access.AccessKeySecret, access.Endpoint)
	if err != nil {
		return nil, err
	}
	a.cdn20180510 = client
	return a, nil
}

func (d *aLiYunCDN) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	domain := d.access.Domain
	if domain == "" {
		domain = certificate.Domain
	}
	resp, err := d.
		cdn20180510.
		SetCdnDomainSSLCertificateWithOptions(&cdn20180510.SetCdnDomainSSLCertificateRequest{
			DomainName:  tea.String(domain),
			CertName:    tea.String(domainUUID(domain)),
			CertType:    tea.String("upload"),
			SSLProtocol: tea.String("on"),
			SSLPub:      tea.String(certificate.Certificate),
			SSLPri:      tea.String(certificate.PrivateKey),
			CertRegion:  tea.String(d.access.Region),
		}, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	d.logs = append(d.logs, AddLog(ALiYunCDN, "Deployment successful", resp))
	return nil
}

func (d *aLiYunCDN) GetLogs() []string {
	return d.logs
}
func (d *aLiYunCDN) createClient(accessKeyId, accessKeySecret, endpoint string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}
