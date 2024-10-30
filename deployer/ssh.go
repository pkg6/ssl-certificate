package deployer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/sftp"
	"github.com/pkg6/ssl-certificate/helper"
	"github.com/pkg6/ssl-certificate/registrations"
	"golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
)

type sshd struct {
	options *Options
	logs    []string
}

type SSHAccess struct {
	Host          string `json:"host" xml:"host" yaml:"host"`
	Username      string `json:"username" xml:"username" yaml:"username"`
	Password      string `json:"password" xml:"password" yaml:"password"`
	Key           string `json:"key" xml:"key" yaml:"key"`
	Port          string `json:"port" xml:"port" yaml:"port"`
	BeforeCommand string `json:"beforeCommand" xml:"beforeCommand" yaml:"beforeCommand"`
	AfterCommand  string `json:"afterCommand" xml:"afterCommand" yaml:"afterCommand"`
	CertPath      string `json:"certPath" xml:"certPath" yaml:"certPath"`
	KeyPath       string `json:"keyPath" xml:"keyPath" yaml:"keyPath"`
}

func NewSSH(options *Options) (IDeployer, error) {
	return &sshd{
		options: options,
		logs:    make([]string, 0),
	}, nil
}
func (d *sshd) GetLogs() []string {
	return d.logs
}
func (d *sshd) Deploy(ctx context.Context, certificate *registrations.Certificate) error {
	access := &SSHAccess{}
	if err := helper.JsonUnmarshal(d.options.Access, access); err != nil {
		return err
	}
	client, err := d.sshClient(access)
	if err != nil {
		return err
	}
	defer client.Close()
	d.logs = append(d.logs, AddLog(SSH, "connection successful", nil))
	if access.BeforeCommand != "" {
		err, stdout, stderr := d.sshExecCommand(client, access.BeforeCommand)
		if err != nil {
			return fmt.Errorf("failed to run before-command: %w, stdout: %s, stderr: %s", err, stdout, stderr)
		}
		d.logs = append(d.logs, AddLog(SSH, "before-command executed successfully", nil))
	}
	// 上传证书
	if err := d.upload(client, certificate.Certificate, access.CertPath); err != nil {
		return fmt.Errorf("failed to upload certificate: %w", err)
	}
	d.logs = append(d.logs, AddLog(SSH, "Successfully upload certificate："+access.CertPath, nil))
	// 上传私钥
	if err := d.upload(client, certificate.PrivateKey, access.KeyPath); err != nil {
		return fmt.Errorf("failed to upload private key: %w", err)
	}
	d.logs = append(d.logs, AddLog(SSH, "Successfully upload private key："+access.KeyPath, nil))
	if access.AfterCommand != "" {
		err, stdout, stderr := d.sshExecCommand(client, access.AfterCommand)
		if err != nil {
			return fmt.Errorf("failed to run command: %w, stdout: %s, stderr: %s", err, stdout, stderr)
		}
		d.logs = append(d.logs, AddLog(SSH, "after-command executed successfully", nil))
	}
	d.logs = append(d.logs, AddLog(Local, "Deployment successful", nil))
	return nil
}
func (d *sshd) upload(client *ssh.Client, content, sshPath string) error {
	sftpCli, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("failed to create sftp client: %w", err)
	}
	defer sftpCli.Close()
	if err := sftpCli.MkdirAll(filepath.Dir(sshPath)); err != nil {
		return fmt.Errorf("failed to create remote directory: %w", err)
	}
	file, err := sftpCli.OpenFile(sshPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return fmt.Errorf("failed to open remote file: %w", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("failed to write to remote file: %w", err)
	}
	return nil
}

func (d *sshd) sshClient(access *SSHAccess) (*ssh.Client, error) {
	var authMethod ssh.AuthMethod
	if access.Key != "" {
		signer, err := ssh.ParsePrivateKey([]byte(access.Key))
		if err != nil {
			return nil, err
		}
		authMethod = ssh.PublicKeys(signer)
	} else {
		authMethod = ssh.Password(access.Password)
	}
	return ssh.Dial("tcp", fmt.Sprintf("%s:%s", access.Host, access.Port), &ssh.ClientConfig{
		User: access.Username,
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
}
func (d *sshd) sshExecCommand(client *ssh.Client, command string) (error, string, string) {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create ssh session: %w", err), "", ""
	}
	defer session.Close()
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	var stderrBuf bytes.Buffer
	session.Stderr = &stderrBuf
	err = session.Run(command)
	return err, stdoutBuf.String(), stderrBuf.String()
}
