package providers

import (
	"github.com/go-acme/lego/v4/providers/http/webroot"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
)

type HTTPAccess struct {
	Path string `json:"path" yaml:"path" xml:"path"`
}

type HTTP struct {
	option *Options
}

func NewHTTP(option *Options) IProvider {
	return &HTTP{
		option: option,
	}
}

func (a *HTTP) Apply() (*registrations.Certificate, error) {
	access := &HTTPAccess{}
	helper.JsonUnmarshal(a.option.Config, access)
	dnsProvider, err := webroot.NewHTTPProvider(access.Path)
	if err != nil {
		return nil, err
	}
	return apply(a.option, dnsProvider)
}
