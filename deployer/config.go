package deployer

const (
	SSH   = "ssh"
	Local = "local"
)

type Config struct {
	Name    string   `json:"name" xml:"name" yaml:"name"`
	Options *Options `json:"options" xml:"options" yaml:"options"`
}

type Options struct {
	Access any `json:"access"`
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
