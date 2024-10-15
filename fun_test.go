package certificate

import (
	"github.com/pkg6/ssl-certificate/registrations"
	"reflect"
	"testing"
)

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
