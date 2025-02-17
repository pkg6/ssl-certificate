package deployer

import (
	"context"
	"fmt"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/registrations"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type LocalAccess struct {
	BeforeCommand string `json:"beforeCommand" xml:"beforeCommand" yaml:"beforeCommand"`
	AfterCommand  string `json:"afterCommand" xml:"afterCommand" yaml:"afterCommand"`
	CertPath      string `json:"certPath" xml:"certPath" yaml:"certPath"`
	KeyPath       string `json:"keyPath" xml:"keyPath" yaml:"keyPath"`
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

func (d *local) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	var (
		wg errgroup.Group
	)
	access := &LocalAccess{}
	if err := pkg.JsonUnmarshal(d.options.Access, access); err != nil {
		return err
	}
	if access.BeforeCommand != "" {
		err := d.execCmd(access.BeforeCommand)
		if err != nil {
			return fmt.Errorf("failed to run before-command: %w", err)
		}
		d.logs = append(d.logs, AddLog(NameLocal, "before-command executed successfully ", nil))
	}
	wg.Go(func() error {
		return d.copyFile(certificate.Certificate, access.CertPath)
	})
	wg.Go(func() error {
		return d.copyFile(certificate.PrivateKey, access.KeyPath)
	})
	if err := wg.Wait(); err != nil {
		d.logs = append(d.logs, AddLog(NameLocal, fmt.Sprintf("Key pair writing failed: %v", err), nil))
		return err
	} else {
		d.logs = append(d.logs, AddLog(NameLocal, "Key pair written successfully", nil))
	}
	if access.AfterCommand != "" {
		if err := d.execCmd(access.AfterCommand); err != nil {
			return fmt.Errorf("failed to run after-command:: %w", err)
		}
		d.logs = append(d.logs, AddLog(NameLocal, "after-command executed successfully", nil))
	}
	d.logs = append(d.logs, AddLog(NameLocal, "Deployment successful", nil))
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
