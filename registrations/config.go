package registrations

import (
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type Config struct {
	Email       string           `json:"email" yaml:"email" xml:"email"`
	Provider    string           `json:"provider" yaml:"provider" xml:"provider"`
	Options     *RegisterOptions `json:"options" yaml:"options" xml:"options"`
	Nameservers string           `json:"nameservers" yaml:"nameservers" xml:"nameservers"`
}

type RegisterOptions struct {
	Kid         string `json:"kid" yaml:"kid" xml:"kid"`
	HmacEncoded string `json:"hmac_encoded" yaml:"hmacEncoded" xml:"hmacEncoded"`
}
type IRegistration interface {
	URL() string
	UserAgent() string
	Resource(lego *lego.Client, opt *RegisterOptions) (*registration.Resource, error)
}
