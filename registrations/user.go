package registrations

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/go-acme/lego/v4/registration"
	"github.com/pkg6/ssl-certificate/helper"
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

func userFileName(email string, regi IRegistration, opt *RegisterOptions) string {
	return helper.HomeDataFile("user-" + email + "-" + helper.MD5String(regi.URL()+regi.UserAgent()))
}

func saveUserData(fileName string, user *User) error {
	privateKeyBytes, err := x509.MarshalECPrivateKey(user.Key.(*ecdsa.PrivateKey))
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %v", err)
	}
	userData := &LegoUserData{
		Email:        user.Email,
		PrivateKey:   string(privateKeyBytes),
		Registration: user.Registration,
	}
	// 将用户数据序列化为 JSON 并写入文件
	data, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}
	return helper.WriteToFile(fileName, data)
}

// LoadUserFromFile loads the user data from a file
func loadUserFromFile(filename string) (*User, error) {
	var userData LegoUserData
	if err := helper.ReadFromFile(filename, &userData); err != nil {
		return nil, err
	}
	privateKey, err := x509.ParseECPrivateKey([]byte(userData.PrivateKey))
	if err != nil {
		return nil, err
	}
	return &User{Email: userData.Email, Key: privateKey, Registration: userData.Registration}, nil
}
