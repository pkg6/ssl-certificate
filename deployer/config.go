package deployer

import "github.com/pkg6/ssl-certificate/pkg"

const (
	NameSSH        = "ssh"
	NameLocal      = "local"
	NameOSS        = "oss"
	NameALiYunCDN  = "aliyunCDN"
	NameALiYunDCDN = "aliyunDCDN"
)

type Config struct {
	Name     string `json:"name" xml:"name" yaml:"name"`
	Deployer IDeployer
	Options  *Options `json:"options" xml:"options" yaml:"options"`
}

type Options struct {
	Access any `json:"access" xml:"Access" yaml:"access"`
}

func (o *Options) JsonUnmarshal(v any) error {
	return pkg.JsonUnmarshal(o, v)
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
