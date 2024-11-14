package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
	"strings"
	"time"
)

func Run(ctx context.Context, config *Config, certificate *registrations.Certificate) ([]string, error) {
	dep, err := NewDeployer(config)
	if err != nil {
		return nil, err
	}
	if err := dep.Deploy(ctx, certificate); err != nil {
		return nil, err
	}
	return dep.GetLogs(), err
}

func NewDeployer(cfg *Config) (IDeployer, error) {
	if cfg.Deployer != nil {
		return cfg.Deployer, nil
	}
	switch cfg.Name {
	case NameSSH:
		return NewSSH(cfg.Options)
	case NameLocal:
		return NewLocal(cfg.Options)
	case NameOSS:
		return NewALiYunOSS(cfg.Options)
	case NameALiYunCDN:
		return NewALiYunCDN(cfg.Options)
	case NameALiYunDCDN:
		return NewALiYunDCDN(cfg.Options)
	default:
		return nil, fmt.Errorf("unknown deployer: %s", cfg.Name)
	}
}

type IDeployer interface {
	Deploy(ctx context.Context, certificate *registrations.Certificate) error
	GetLogs() []string
}

func domainUUID(domain string) string {
	return fmt.Sprintf("%s-%s", domain, time.Now().Format("20060102150405"))
}

func AddLog(deployer string, log string, v any) string {
	return strings.Join([]string{"【" + deployer + "】:", log, pkg.JsonMarshal(v)}, " ")
}
