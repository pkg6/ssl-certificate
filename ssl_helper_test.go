package certificate

import (
	"context"
	"github.com/pkg6/ssl-certificate/deployer"
	"github.com/pkg6/ssl-certificate/registrations"
	"reflect"
	"testing"
)

func TestDeployer(t *testing.T) {
	type args struct {
		ctx         context.Context
		config      *deployer.Config
		certificate *registrations.Certificate
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Deployer(tt.args.ctx, tt.args.config, tt.args.certificate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Deployer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deployer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSSLCertificate(t *testing.T) {
	type args struct {
		email          string
		domain         []string
		provider       string
		providerConfig any
	}
	tests := []struct {
		name    string
		args    args
		want    *registrations.Certificate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SSLCertificate(tt.args.email, tt.args.domain, tt.args.provider, tt.args.providerConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("SSLCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SSLCertificate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSSLCertificateByConfig(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		want    *registrations.Certificate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SSLCertificateByConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("SSLCertificateByConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SSLCertificateByConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSSLCertificateDeployer(t *testing.T) {
	type args struct {
		ctx      context.Context
		cfg      *Config
		deployer *deployer.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SSLCertificateDeployer(tt.args.ctx, tt.args.cfg, tt.args.deployer); (err != nil) != tt.wantErr {
				t.Errorf("SSLCertificateDeployer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
