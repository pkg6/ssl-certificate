package registrations

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/http/webroot"
	"github.com/pkg6/ssl-certificate/helper"
	"strings"
)

const (
	LetsencryptSSL = "letsencrypt"
	ZeroSSL        = "zerossl"
)

var registrations = map[string]IRegistration{
	LetsencryptSSL: &letsencryptRegistration{},
	ZeroSSL:        &zeroRegistration{},
}

func GetRegistration(name string) (IRegistration, error) {
	if reg, ok := registrations[name]; ok {
		return reg, nil
	}
	return nil, fmt.Errorf("registrations %s not found", name)
}

func LegoClient(email string, regi IRegistration, opt *RegisterOptions) (*User, *lego.Client, error) {
	if email == "" {
		email = helper.UUIDEmail()
	}
	data := NewData(email, regi, opt)
	var (
		user *User
	)
	user, _ = data.LoadUser()
	if user == nil {
		user = &User{Email: email}
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return user, nil, err
		}
		user.Key = privateKey
	}
	config := lego.NewConfig(user)
	config.CADirURL = regi.URL()
	config.Certificate.KeyType = certcrypto.RSA2048
	if userAgent := regi.UserAgent(); userAgent != "" {
		config.UserAgent = userAgent
	}
	client, err := lego.NewClient(config)
	if err != nil {
		return user, nil, err
	}
	user.Registration, err = regi.Resource(client, opt)
	if err != nil {
		return user, nil, err
	}
	_ = data.SaveUser(user)
	return user, client, nil
}

func RegistrationByProvider(provider challenge.Provider, registration *Config, domains []string) (*Certificate, error) {
	if registration.Provider == "" {
		registration.Provider = LetsencryptSSL
	}
	if registration.Options == nil {
		registration.Options = &RegisterOptions{}
	}
	regi, err := GetRegistration(registration.Provider)
	if err != nil {
		return nil, err
	}
	user, legoClient, err := LegoClient(registration.Email, regi, registration.Options)
	if err != nil {
		return nil, err
	}
	var (
		cert *certificate.Resource
	)
	if prov, ok := provider.(*webroot.HTTPProvider); ok {
		cert, err = UseHTTPObtain(legoClient, prov, domains)
	} else {
		cert, err = UseDNSObtain(legoClient, provider, domains, strings.Split(registration.Nameservers, ","))
	}
	if err != nil {
		return nil, err
	}
	return NewCertificateByResource(user, cert), nil
}

func UseDNSObtain(legoClient *lego.Client, provider challenge.Provider, domains []string, nameservers []string) (*certificate.Resource, error) {
	challengeOptions := make([]dns01.ChallengeOption, 0)
	if len(nameservers) > 0 {
		challengeOptions = append(challengeOptions, dns01.AddRecursiveNameservers(nameservers))
	}
	_ = legoClient.Challenge.SetDNS01Provider(provider, challengeOptions...)
	return legoClient.Certificate.Obtain(certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	})
}

func UseHTTPObtain(legoClient *lego.Client, provider challenge.Provider, domains []string) (*certificate.Resource, error) {
	_ = legoClient.Challenge.SetHTTP01Provider(provider)
	return legoClient.Certificate.Obtain(certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	})
}
