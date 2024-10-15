package helper

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net"
)

func ID() string {
	faces, _ := net.Interfaces()
	var s string
	for _, address := range faces {
		if macAddr := address.HardwareAddr.String(); macAddr != "" {
			s += macAddr
		}
	}
	return MD5String(s)
}

func MD5String(s string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}
