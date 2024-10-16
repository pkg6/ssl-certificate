package deployer

import (
	"context"
	"encoding/json"
	"fmt"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg6/ssl-certificate/registrations"
	"time"
)

type aLiYunCDN struct {
	cdn20180510 *cdn20180510.Client
	options     *Options
	access      *ALiYunCDNAccess
	logs        []string
}
type ALiYunCDNAccess struct {
	Domain          string `json:"domain" xml:"domain" yaml:"domain"`
	AccessKeyId     string `json:"accessKeyId" xml:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" xml:"accessKeySecret" yaml:"accessKeySecret"`
	Region          string `json:"region" xml:"region" yaml:"region"`
}

func NewALiYunCDN(options *Options) (IDeployer, error) {
	access := &ALiYunCDNAccess{}
	_ = options.JsonUnmarshal(access)
	if access.Region == "" {
		access.Region = "cn-hangzhou"
	}
	a := &aLiYunCDN{
		options: options,
		access:  access,
	}
	client, err := a.createClient(access.AccessKeyId, access.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	a.cdn20180510 = client
	return a, nil
}

func (a aLiYunCDN) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	certName := fmt.Sprintf("%s-%s", a.access.Domain, time.Now().Format("20060102150405"))
	setCdnDomainSSLCertificateRequest := &cdn20180510.SetCdnDomainSSLCertificateRequest{
		DomainName:  tea.String(a.access.Domain),
		CertName:    tea.String(certName),
		CertType:    tea.String("upload"),
		SSLProtocol: tea.String("on"),
		SSLPub:      tea.String(certificate.Certificate),
		SSLPri:      tea.String(certificate.PrivateKey),
		CertRegion:  tea.String(a.access.Region),
	}
	runtime := &util.RuntimeOptions{}
	resp, err := a.cdn20180510.SetCdnDomainSSLCertificateWithOptions(setCdnDomainSSLCertificateRequest, runtime)
	if err != nil {
		return err
	}
	respByte, _ := json.Marshal(resp)
	a.logs = append(a.logs, "【ALiYun CDN】"+string(respByte))
	return nil
}

func (a *aLiYunCDN) GetLogs() []string {
	return a.logs
}
func (a *aLiYunCDN) createClient(accessKeyId, accessKeySecret string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}
