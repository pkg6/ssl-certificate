package main

import (
	"context"
	"flag"
	"github.com/go-acme/lego/v4/log"
	certificate "github.com/pkg6/ssl-certificate"
	"github.com/pkg6/ssl-certificate/helper"
)

var (
	cfg  *certificate.DomainsDeploysConfig
	file string
	err  error
)

func init() {
	flag.StringVar(&file, "f", "config.json", "User-defined configuration files")
}
func main() {
	flag.Parse()
	err = helper.JsonFileUnmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("load config err=%v", err)
		return
	}
	for _, domain := range cfg.Domains {
		ssl, err := certificate.SSLCertificateByConfig(domain.Certificate)
		if err != nil {
			log.Fatalf("load ssl err=%v", err)
			return
		}
		logs, err := certificate.Deployer(domain.DeployerConfig(cfg.Deploys), context.Background(), ssl)
		if err != nil {
			log.Fatalf("deploy err=%v", err)
			return
		}
		for _, l := range logs {
			log.Println(l)
		}
	}
}
