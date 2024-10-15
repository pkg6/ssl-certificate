package main

import (
	"fmt"
	certificate "github.com/pkg6/ssl-certificate"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

func main() {
	config := &certificate.Config{
		Domains: []string{"ssl.zhiqiang.wang"},
		Provider: &providers.Config{
			Name: "aliyun",
			Config: &providers.AliYunAccess{
				AccessKeyId:     "****************",
				AccessKeySecret: "****************",
			},
			//Name: "http",
			//Config: &providers.HTTPAccess{
			//	Path: "/data/wwwroot/ssl.zhiqiang.wang",
			//},
		},
		Registration: &registrations.Config{
			Provider: registrations.LetsencryptSSL,
		},
	}
	ssl, err := certificate.SSLCertificateByConfig(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("证书：%s\n", ssl.Certificate)
	fmt.Printf("密钥：%s\n", ssl.PrivateKey)
	fmt.Printf("CA证书：%s\n", ssl.IssuerCertificate)
}
