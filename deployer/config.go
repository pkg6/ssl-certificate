package deployer

import "github.com/pkg6/ssl-certificate/helper"

const (
	SSH        = "ssh"
	Local      = "local"
	OSS        = "oss"
	ALiYunCDN  = "aliyunCDN"
	ALiYunDCDN = "aliyunDCDN"
	TencentCDN = "tencentCDN"
)

type Config struct {
	Name    string   `json:"name" xml:"name" yaml:"name"`
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

type Options struct {
	Access any `json:"access" xml:"Access" yaml:"access"`
}

func (o *Options) JsonUnmarshal(v any) error {
	return helper.JsonUnmarshal(o, v)
}

func MapNameAny(name string, access any) *Config {
	if access == nil {
		access = ""
	}
	return &Config{
		Name:    name,
		Options: &Options{Access: access},
	}
}
