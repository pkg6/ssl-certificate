package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-acme/lego/v4/log"
	certificate "github.com/pkg6/ssl-certificate"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/providers"
	"github.com/pkg6/ssl-certificate/registrations"
	"path"
)

var (
	domain  string
	webroot string
	sslPath string
)

func init() {
	flag.StringVar(&domain, "domain", "", "Need to generate SSL domain names")
	flag.StringVar(&webroot, "webroot", "", "Directory for domain deployment")
	flag.StringVar(&sslPath, "path", "/etc/nginx/ssl/", "Directory for storing certificates")
}
func main() {
	flag.Parse()
	if domain == "" {
		log.Fatal("Need to set --domain")
	}
	if webroot == "" {
		log.Fatal("Need to set --webroot")
	}
	if sslPath == "" {
		log.Fatal("Need to set --path")
	}

	ssl, err := certificate.SSLCertificateByConfig(&certificate.Config{
		Domains:      []string{domain},
		Registration: &registrations.Config{},
		Provider: &providers.Config{
			Name:   providers.Http,
			Config: &providers.HTTPAccess{Path: webroot},
		},
	})
	if err != nil {
		log.Fatal("Failed to generate SSL domain names")
	}
	logs, err := certificate.Deployer(&deployer.Config{
		Name: deployer.Local,
		Options: &deployer.Options{
			Access: &deployer.LocalAccess{
				AfterCommand: "service nginx restart",
				CertPath:     path.Join(sslPath, fmt.Sprintf("%s.crt", domain)),
				KeyPath:      path.Join(sslPath, fmt.Sprintf("%s.key", domain)),
			},
		},
	}, context.Background(), ssl)
	if err != nil {
		log.Fatalf("deploy err=%v", err)
		return
	}
	for _, l := range logs {
		log.Println(l)
	}
}
