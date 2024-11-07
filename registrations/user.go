package registrations

import (
	"crypto"
	"github.com/go-acme/lego/v4/registration"
)

type User struct {
	Email        string                 `json:"email"`
	Key          crypto.PrivateKey      `json:"key"`
	Registration *registration.Resource `json:"registration"`
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

type LegoUserData struct {
	Email        string                 `json:"email"`
	PrivateKey   string                 `json:"private_key"`
	Registration *registration.Resource `json:"registration"`
}
