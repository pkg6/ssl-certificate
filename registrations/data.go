package registrations

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/pkg6/ssl-certificate/pkg"
	"github.com/pkg6/ssl-certificate/pkg/data"
	"net/url"
	"path"
)

type Data struct {
	filename string
	data     *data.LocalData[*LegoUserData]
}

func NewData(email string, regi IRegistration, opt *RegisterOptions) *Data {
	urlP, _ := url.Parse(regi.URL())
	optByte, _ := json.Marshal(opt)
	userPath := path.Join("user", urlP.Host, fmt.Sprintf("%s-%s", email, pkg.MD5String(string(optByte))))
	filName := pkg.HomeDataFile(userPath)
	return &Data{data: data.NewLocalData[*LegoUserData](filName)}
}

func (d *Data) SaveUser(user *User) error {
	privateKeyBytes, err := x509.MarshalECPrivateKey(user.Key.(*ecdsa.PrivateKey))
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %v", err)
	}
	userData := &LegoUserData{Email: user.Email, PrivateKey: string(privateKeyBytes), Registration: user.Registration}
	return d.data.Save(userData)
}

func (d *Data) LoadUser() (*User, error) {
	userData, err := d.data.Load()
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParseECPrivateKey([]byte(userData.PrivateKey))
	if err != nil {
		return nil, err
	}
	return &User{Email: userData.Email, Key: privateKey, Registration: userData.Registration}, nil
}
