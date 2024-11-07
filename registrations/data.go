package registrations

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/pkg6/ssl-certificate/helper"
	"net/url"
	"path"
)

type Data struct {
	filename string
}

func NewData(email string, regi IRegistration, opt *RegisterOptions) *Data {
	urlP, _ := url.Parse(regi.URL())
	optByte, _ := json.Marshal(opt)
	userPath := path.Join("user", urlP.Host, fmt.Sprintf("%s-%s", email, helper.MD5String(string(optByte))))
	return &Data{filename: helper.HomeDataFile(userPath)}
}

func (d *Data) SaveUser(user *User) error {
	privateKeyBytes, err := x509.MarshalECPrivateKey(user.Key.(*ecdsa.PrivateKey))
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %v", err)
	}
	userData := &LegoUserData{Email: user.Email, PrivateKey: string(privateKeyBytes), Registration: user.Registration}
	// 将用户数据序列化为 JSON 并写入文件
	data, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}
	return helper.WriteToFile(d.filename, data)
}

func (d *Data) LoadUser() (*User, error) {
	var userData LegoUserData
	if err := helper.ReadFromFile(d.filename, &userData); err != nil {
		return nil, err
	}
	privateKey, err := x509.ParseECPrivateKey([]byte(userData.PrivateKey))
	if err != nil {
		return nil, err
	}
	return &User{Email: userData.Email, Key: privateKey, Registration: userData.Registration}, nil
}
