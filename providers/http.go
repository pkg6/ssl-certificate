package providers

import (
	"github.com/go-acme/lego/v4/providers/http/webroot"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
)

type HTTPAccess struct {
	Path string `json:"path" yaml:"path" xml:"path"`
}

type http struct {
	option *Options
}

func NewHTTP(option *Options) IProvider {
	return &http{
		option: option,
	}
}

func (a *http) Apply() (*registrations.Certificate, error) {
	access := &HTTPAccess{}
	helper.JsonUnmarshal(a.option.Config, access)
	dnsProvider, err := webroot.NewHTTPProvider(access.Path)
	if err != nil {
		return nil, err
	}
	return apply(a.option, dnsProvider)
}
