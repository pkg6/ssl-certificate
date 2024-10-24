package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"strings"
	"time"
)

func NewDeployer(cfg *Config) (IDeployer, error) {
	switch cfg.Name {
	case SSH:
		return NewSSH(cfg.Options)
	case Local:
		return NewLocal(cfg.Options)
	case OSS:
		return NewALiYunOSS(cfg.Options)
	case ALiYunCDN:
		return NewALiYunCDN(cfg.Options)
	case ALiYunDCDN:
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
	return strings.Join([]string{"【" + deployer + "】:", log, helper.JsonMarshal(v)}, " ")
}
