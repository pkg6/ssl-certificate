package helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/shirou/gopsutil/v4/host"
	"io"
	"net"
	"runtime"
)

func UUID() string {
	var uuid string
	info, _ := host.Info()
	if info.HostID != "" {
		uuid = info.HostID
	}
	if uuid == "" {
		faces, _ := net.Interfaces()
		for _, address := range faces {
			if macAddr := address.HardwareAddr.String(); macAddr != "" {
				uuid += macAddr
			}
		}
	}
	hash := md5.New()
	_, _ = io.WriteString(hash, uuid)
	return hex.EncodeToString(hash.Sum(nil))

}

func UUIDEmail() string {
	return UUID() + "@" + runtime.GOOS + ".com"
}

func MD5String(s string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}
