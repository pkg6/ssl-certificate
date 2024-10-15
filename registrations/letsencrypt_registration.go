package registrations

import (
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type letsencryptRegistration struct {
}

func (l letsencryptRegistration) Resource(lego *lego.Client, opt *RegisterOptions) (*registration.Resource, error) {
	return lego.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
}
func (l letsencryptRegistration) URL() string {
	return lego.LEDirectoryProduction
}

func (l letsencryptRegistration) UserAgent() string {
	return ""
}
