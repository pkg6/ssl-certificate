package providers

import (
	"fmt"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/pkg6/ssl-certificate/registrations"
)

type IProvider interface {
	Apply() (*registrations.Certificate, error)
}

func apply(options *Options, provider challenge.Provider) (*registrations.Certificate, error) {
	return registrations.RegistrationByProvider(provider, options.Registration, options.Domains)
}

func NewProvider(cfg *Config, registration *registrations.Config, domains []string) (IProvider, error) {
	if cfg.Provider != nil {
		return cfg.Provider, nil
	}
	option := &Options{Domains: domains, Config: cfg.Config, Registration: registration}
	switch cfg.Name {
	case Aliyun:
		return NewAliyun(option), nil
	case Tencent:
		return NewTencent(option), nil
	case Huaweicloud:
		return NewHuaweiCloud(option), nil
	case Cloudflare:
		return NewCloudflare(option), nil
	case Godaddy:
		return NewGodaddy(option), nil
	case Http:
		return NewHTTP(option), nil
	case AWS:
		return NewAws(option), nil
	case Powerdns:
		return NewPowerdns(option), nil
	default:
		return nil, fmt.Errorf("unknown %s config provider", cfg.Name)
	}
}
