package helper

import (
	"fmt"
	"os"
)

const (
	TimeOutKey     = "SSL-CERTIFICATE-TIMEOUT"
	defaultTimeOut = 60
)

func Setenv(key, value string) error {
	val := os.Getenv(key)
	if val == "" {
		return os.Setenv(key, value)
	}
	return nil
}

func SetTimeOut(key string) error {
	val := os.Getenv(TimeOutKey)
	if val == "" {
		val = fmt.Sprintf("%d", defaultTimeOut)
	}
	return Setenv(key, val)
}
