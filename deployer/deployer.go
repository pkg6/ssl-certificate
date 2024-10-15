package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/registrations"
)

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

func NewDeployer(cfg *Config) (IDeployer, error) {
	switch cfg.Name {
	case SSH:
		return NewSSH(cfg.Options)
	case Local:
		return NewLocal(cfg.Options)
	default:
		return nil, fmt.Errorf("unknown deployer: %s", cfg.Name)
	}
}

type IDeployer interface {
	Deploy(ctx context.Context, certificate registrations.Certificate) error
	GetLogs() []string
}
