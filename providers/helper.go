package providers

import (
	"fmt"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
	"strconv"
	"strings"
)

type IProvider interface {
	Apply() (*registrations.Certificate, error)
}

func Apply(options *Options, provider challenge.Provider) (*registrations.Certificate, error) {
	//https://github.com/go-acme/lego/issues/1867
	for _, domain := range options.Domains {
		if strings.HasPrefix(domain, "*") {
			_ = pkg.Setenv("LEGO_DISABLE_CNAME_SUPPORT", strconv.FormatBool(true))
			break
		}
	}
	return registrations.RegistrationByProvider(provider, options.Registration, options.Domains)
}

func NewProvider(cfg *Config, registration *registrations.Config, domains []string) (IProvider, error) {
	if cfg.Provider != nil {
		return cfg.Provider, nil
	}
	option := &Options{Domains: domains, Config: cfg.Config, Registration: registration}
	switch cfg.Name {
	case NameALiYun:
		return NewALiYun(option), nil
	case NameTencent:
		return NewTencent(option), nil
	case NameHuawei:
		return NewHuaweiCloud(option), nil
	case NameCloudflare:
		return NewCloudflare(option), nil
	case NameGodaddy:
		return NewGodaddy(option), nil
	case NameHTTP:
		return NewHTTP(option), nil
	case NameAWS:
		return NewAws(option), nil
	case NamePowerdns:
		return NewPowerdns(option), nil
	case NameNamesilo:
		return NewNameSilo(option), nil
	default:
		return nil, fmt.Errorf("unknown %s config provider", cfg.Name)
	}
}
