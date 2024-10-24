package deployer

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dcdn20180115 "github.com/alibabacloud-go/dcdn-20180115/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg6/ssl-certificate/registrations"
	"strings"
)

type ALiYunDCDNAccess struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Domain          string `json:"domain"`
	Region          string `json:"region"`
	Endpoint        string `json:"endpoint"`
}
type aLiYunDCDN struct {
	options *Options
	logs    []string
	access  *ALiYunDCDNAccess
	client  *dcdn20180115.Client
}

func NewALiYunDCDN(options *Options) (IDeployer, error) {
	access := &ALiYunDCDNAccess{}
	_ = options.JsonUnmarshal(access)
	if access.Region == "" {
		access.Region = "cn-hangzhou"
	}
	if access.Endpoint == "" {
		access.Endpoint = "dcdn.aliyuncs.com"
	}
	d := &aLiYunDCDN{
		options: options,
		logs:    make([]string, 0),
		access:  access,
	}
	client, err := d.createClient(access.AccessKeyId, access.AccessKeySecret, access.Endpoint)
	if err != nil {
		return nil, err
	}
	d.client = client
	return d, nil
}
func (d *aLiYunDCDN) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	domain := d.access.Domain
	if domain == "" {
		domain = certificate.Domain
	}
	if strings.HasPrefix(domain, "*") {
		domain = strings.TrimPrefix(domain, "*")
	}
	resp, err := d.client.SetDcdnDomainSSLCertificateWithOptions(&dcdn20180115.SetDcdnDomainSSLCertificateRequest{
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
	d.logs = append(d.logs, AddLog(ALiYunDCDN, "Deployment successful", resp))
	return nil
}

func (d *aLiYunDCDN) GetLogs() []string {
	return d.logs
}
func (d *aLiYunDCDN) createClient(accessKeyId, accessKeySecret, endpoint string) (_result *dcdn20180115.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	_result = &dcdn20180115.Client{}
	_result, _err = dcdn20180115.NewClient(config)
	return _result, _err
}
