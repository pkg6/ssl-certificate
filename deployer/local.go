package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type LocalAccess struct {
	BeforeCommand string `json:"beforeCommand"`
	AfterCommand  string `json:"afterCommand"`
	CertPath      string `json:"certPath"`
	KeyPath       string `json:"keyPath"`
}

type local struct {
	options *Options
	logs    []string
}

func NewLocal(option *Options) (IDeployer, error) {
	return &local{
		options: option,
		logs:    make([]string, 0),
	}, nil
}
func (d *local) GetLogs() []string {
	return d.logs
}

func (d *local) Deploy(ctx context.Context, certificate registrations.Certificate) error {
	access := &LocalAccess{}
	if err := helper.JsonUnmarshal(d.options.Access, access); err != nil {
		return err
	}
	if access.BeforeCommand != "" {
		err := d.execCmd(access.BeforeCommand)
		if err != nil {
			return fmt.Errorf("failed to run before-command: %w", err)
		}
		d.logs = append(d.logs, "【local】 before-command executed successfully")
	}
	// 复制文件
	if err := d.copyFile(certificate.Certificate, access.CertPath); err != nil {
		return fmt.Errorf("copy certificate failed: %w", err)
	}
	d.logs = append(d.logs, "【local】 certificate upload successful")
	if err := d.copyFile(certificate.PrivateKey, access.KeyPath); err != nil {
		return fmt.Errorf("copy private key failed: %w", err)
	}
	d.logs = append(d.logs, "【local】 successfully uploaded private key")
	if access.AfterCommand != "" {
		if err := d.execCmd(access.AfterCommand); err != nil {
			return fmt.Errorf("failed to run after-command:: %w", err)
		}
		d.logs = append(d.logs, "【local】 after-command executed successfully")
	}
	return nil
}
func (d local) execCmd(command string) error {
	// 执行命令
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to carry out command: %w", err)
	}
	return nil
}
func (d local) copyFile(content string, path string) error {
	dir := filepath.Dir(path)
	// 如果目录不存在，创建目录
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	// 创建或打开文件
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	// 写入内容到文件
	_, err = file.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("fail to write to file: %w", err)
	}
	return nil
}
