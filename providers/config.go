package providers

import "github.com/pkg6/ssl-certificate/registrations"

const (
	Aliyun      = "aliyun"
	Tencent     = "tencent"
	Huaweicloud = "huaweicloud"
	Cloudflare  = "cloudflare"
	Godaddy     = "godaddy"
	Http        = "http"
)

type Config struct {
	Name   string `json:"name" xml:"name" yaml:"name"`
	Config any    `json:"config" xml:"config" yaml:"config"`
}

type Options struct {
	Domains      []string              `json:"domain"`
	Config       any                   `json:"config"`
	Registration *registrations.Config `json:"registration"`
}
