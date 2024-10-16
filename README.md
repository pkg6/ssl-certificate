# ssl-certificate

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
	fmt.Printf("证书：%s\n", ssl.Certificate)
	fmt.Printf("密钥：%s\n", ssl.PrivateKey)
	fmt.Printf("CA证书：%s\n", ssl.IssuerCertificate)
}
~~~

nginx ssl 配置

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

