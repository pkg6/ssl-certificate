package providers

import "github.com/pkg6/ssl-certificate/registrations"

const (
	NameALiYun     = "aliyun"
	NameTencent    = "tencent"
	NameHuawei     = "huawei"
	NameCloudflare = "cloudflare"
	NameGodaddy    = "godaddy"
	NameHttp       = "http"
	NameAWS        = "aws"
	NamePowerdns   = "powerdns"
)

type Config struct {
	Name     string `json:"name" xml:"name" yaml:"name"`
	Provider IProvider
	Config   any `json:"config" xml:"config" yaml:"config"`
}

type Options struct {
	Domains      []string              `json:"domain" xml:"domains" yaml:"domains"`
	Config       any                   `json:"config" xml:"config" yaml:"config"`
	Registration *registrations.Config `json:"registration" xml:"registration" yaml:"registration"`
}
