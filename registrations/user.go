package registrations

import (
	"crypto"
	"github.com/go-acme/lego/v4/registration"
)

type User struct {
	Email        string
	key          crypto.PrivateKey
	Registration *registration.Resource
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}
