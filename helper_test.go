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
		config      *deployer.Config
		ctx         context.Context
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
			got, err := Deployer(tt.args.config, tt.args.ctx, tt.args.certificate)
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

func TestDomainCertificates(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    []*CertificateInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DomainCertificates(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainCertificates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainCertificates() got = %v, want %v", got, tt.want)
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
