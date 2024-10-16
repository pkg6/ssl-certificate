## Using the Command Line

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
	fmt.Printf("certificate：%s\n", ssl.Certificate)
	fmt.Printf("privateKey：%s\n", ssl.PrivateKey)
	fmt.Printf("IssuerCertificate：%s\n", ssl.IssuerCertificate)

	//Obtain domain certificate information
	domainCertificates, err := certificate.DomainCertificates("ssl.zhiqiang.wang")
	if err != nil {
		panic(err)
	}
	domainCertificate := domainCertificates[0]
	fmt.Printf("certificate NotBefore：%s\n", domainCertificate.NotBefore)
	fmt.Printf("certificate NotAfter：%s\n", domainCertificate.NotAfter)
}
~~~

