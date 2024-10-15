package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/registrations"
)

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
	Deploy(ctx context.Context, certificate *registrations.Certificate) error
	GetLogs() []string
}
