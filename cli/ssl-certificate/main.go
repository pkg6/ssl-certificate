package main

import (
	"context"
	"flag"
	"github.com/go-acme/lego/v4/log"
	certificate "github.com/pkg6/ssl-certificate"
	"github.com/pkg6/ssl-certificate/pkg"
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
	err = pkg.JsonFileUnmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("load config err=%v", err)
		return
	}
	for _, domain := range cfg.Domains {
		err = certificate.SSLCertificateDeployer(context.Background(), domain.Certificate, domain.DeployerConfig(cfg.Deploys))
		if err != nil {
			log.Fatalf("deploy err=%v", err)
			continue
		}
	}
}
