## Using the Command Line

### sh install or uninstall

~~~
curl -sSL https://raw.githubusercontent.com/pkg6/sh/main/install-pkg.sh | bash -s jq

// install
curl -sSL https://raw.githubusercontent.com/pkg6/ssl-certificate/main/install.sh | bash
// uinstall
curl -sSL https://raw.githubusercontent.com/pkg6/ssl-certificate/main/uninstall.sh | bash
~~~

### ssl-certificate-local

~~~
go install github.com/pkg6/ssl-certificate/cli/ssl-certificate-local@latest
ssl-certificate-local --domain=ssl.zhiqiang.wang --webroot=/data/wwwroot/ssl.zhiqiang.wang --path=/etc/nginx/ssl/ --command="servcie nginx reload"
~~~

## Nginx SSL configuration (partial)

~~~
listen 443 ssl;
ssl_certificate /etc/nginx/ssl/ssl.zhiqiang.wang.cer;
ssl_certificate_key /etc/nginx/ssl/ssl.zhiqiang.wang.key;
ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
ssl_prefer_server_ciphers on;
ssl_session_cache shared:SSL:10m;
ssl_session_timeout 10m;
~~~

## Using SSL certificate for function calls

### download

~~~
go get github.com/pkg6/ssl-certificate
~~~

### Case code

~~~
package main

import (
	"context"
	"fmt"
	certificate "github.com/pkg6/ssl-certificate"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
)

func main() {
	//---------------------generate start-------------------------------
	config := &certificate.Config{
		Domains: []string{"ssl.zhiqiang.wang"},
		Provider: &providers.Config{
			//Name: providers.NameALiYun,
			//Config: &providers.AliYunAccess{
			//	AccessKeyId:     "****************",
			//	AccessKeySecret: "****************",
			//},
			Name: providers.NameHTTP,
			Config: &providers.HTTPAccess{
				Path: "/data/wwwroot/ssl.zhiqiang.wang",
			},
		},
		Registration: &registrations.Config{
			Provider: registrations.LetsencryptSSL,
		},
	}
	ssl, err := certificate.SSLCertificateByConfig(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("certificate：%s\n", ssl.Certificate)
	fmt.Printf("privateKey：%s\n", ssl.PrivateKey)
	fmt.Printf("IssuerCertificate：%s\n", ssl.IssuerCertificate)
	//---------------------generate end-------------------------------

	//---------------------deployer start-------------------------------
	certificate.Deployer(context.Background(), &deployer.Config{
		//Name: deployer.Local,
		//Options: &deployer.Options{Access: deployer.LocalAccess{
		//	CertPath:     "/etc/nginx/etc/ssl.zhiqiang.wang.cer",
		//	KeyPath:      "/etc/nginx/etc/ssl.zhiqiang.wang.key",
		//	AfterCommand: "service nginx reload",
		//}},
		//
		//Name: deployer.SSH,
		//Options: &deployer.Options{Access: deployer.SSHAccess{
		//	Host:         "127.0.0.1",
		//	Username:     "ubuntu",
		//	Password:     "123456",
		//	CertPath:     "/etc/nginx/etc/ssl.zhiqiang.wang.cer",
		//	KeyPath:      "/etc/nginx/etc/ssl.zhiqiang.wang.key",
		//	AfterCommand: "service nginx reload",
		//}},
		//Name: deployer.OSS,
		//Options: &deployer.Options{Access: deployer.ALiYunOSSAccess{
		//	AccessKeyId:     "***********************",
		//	AccessKeySecret: "***********************",
		//	//https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
		//	Endpoint: "oss-cn-hangzhou.aliyuncs.com",
		//	Bucket:   "test",
		//	Domain:   "ssl.zhiqiang.wang",
		//}},
		//Name: deployer.ALiYunCDN,
		//Options: &deployer.Options{Access: deployer.ALiYunCDNAccess{
		//	AccessKeyId:     "***********************",
		//	AccessKeySecret: "***********************",
		//	Endpoint:        "cdn.aliyuncs.com",
		//	Region:          "cn-hangzhou",
		//	Domain:          "ssl.zhiqiang.wang",
		//}},
		//Name: deployer.ALiYunDCDN,
		Options: &deployer.Options{Access: deployer.ALiYunDCDNAccess{
			AccessKeyId:     "***********************",
			AccessKeySecret: "***********************",
			Endpoint:        "dcdn.aliyuncs.com",
			Region:          "cn-hangzhou",
			Domain:          "ssl.zhiqiang.wang",
		}},
	}, ssl)
	//---------------------deployer end-------------------------------

	//---------------------Certificate Information start-------------------------------
	//Obtain certificate information through domain access
	domainCertificates, err := certificate.SSLCertificateInfoByTCP("ssl.zhiqiang.wang")
	if err != nil {
		panic(err)
	}
	domainCertificate := domainCertificates[0]
	fmt.Printf("certificate NotBefore：%s\n", domainCertificate.NotBefore)
	fmt.Printf("certificate NotAfter：%s\n", domainCertificate.NotAfter)

	//Obtain certificate information through the content of the certificate file
	domainCertificate2, err := certificate.SSLCertificateInfoByCer([]byte(ssl.Certificate))
	if err != nil {
		panic(err)
	}
	fmt.Printf("certificate NotBefore：%s\n", domainCertificate2.NotBefore)
	fmt.Printf("certificate NotAfter：%s\n", domainCertificate2.NotAfter)
	//---------------------Certificate Information start-------------------------------
}
~~~

