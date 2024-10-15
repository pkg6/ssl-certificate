package registrations

import (
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type zeroRegistration struct {
}

func (z zeroRegistration) Resource(lego *lego.Client, opt *RegisterOptions) (*registration.Resource, error) {
	return lego.Registration.RegisterWithExternalAccountBinding(registration.RegisterEABOptions{
		TermsOfServiceAgreed: true,
		Kid:                  opt.Kid,
		HmacEncoded:          opt.HmacEncoded,
	})
}
func (z zeroRegistration) URL() string {
	return "https://acme.zerossl.com/v2/DV90"
}
func (z zeroRegistration) UserAgent() string {
	return ""
}
