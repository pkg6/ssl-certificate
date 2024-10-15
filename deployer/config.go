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
	Domain string `json:"domain"`
	Access any    `json:"access"`
}
