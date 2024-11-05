package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const blockType = "PUBLIC KEY"

type SSLCertificateInfo struct {
	Subject            string                  `json:"subject" xml:"subject"`
	Issuer             string                  `json:"issuer" xml:"issuer"`
	NotBefore          time.Time               `json:"not_before" xml:"NotBefore"`
	NotAfter           time.Time               `json:"not_after" xml:"NotAfter"`
	PublicKeyAlgorithm x509.PublicKeyAlgorithm `json:"public_key_algorithm" xml:"PublicKeyAlgorithm"`
	Version            int                     `json:"version" xml:"version"`
	PublicKey          string                  `json:"public_key" xml:"PublicKey"`
}

func x509Certificate(certificate *x509.Certificate) *SSLCertificateInfo {
	info := &SSLCertificateInfo{
		Subject:            certificate.Subject.String(),
		Issuer:             certificate.Issuer.String(),
		NotBefore:          certificate.NotBefore,
		NotAfter:           certificate.NotAfter,
		PublicKeyAlgorithm: certificate.PublicKeyAlgorithm,
		Version:            certificate.Version,
	}
	pubKey, _ := x509.MarshalPKIXPublicKey(certificate.PublicKey)
	pemBlock := &pem.Block{Type: blockType, Bytes: pubKey}
	info.PublicKey = string(pem.EncodeToMemory(pemBlock))
	return info
}

// SSLCertificateInfoByTCP
// Obtain domain certificate information
// certPEM, err := os.ReadFile("ssl.zhiqiang.wang.cer")
// cer, err := certificate.SSLCertificateInfoByTCP("ssl.zhiqiang.wang")
func SSLCertificateInfoByTCP(domain string) ([]*SSLCertificateInfo, error) {
	if strings.HasPrefix(domain, "https://") || strings.HasPrefix(domain, "http://") {
		parse, err := url.Parse(domain)
		if err != nil {
			return nil, err
		}
		domain = parse.Host
	}
	dial, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), nil)
	if err != nil {
		return nil, err
	}
	state := dial.ConnectionState()
	var infos []*SSLCertificateInfo
	for _, certificate := range state.PeerCertificates {
		infos = append(infos, x509Certificate(certificate))
	}
	return infos, nil
}

// SSLCertificateInfoByCer
// Obtain certificate information through the car file
// certPEM, err := os.ReadFile("ssl.zhiqiang.wang.cer")
// cer, err := certificate.SSLCertificateInfoByCer(certPEM)
func SSLCertificateInfoByCer(cer []byte) (*SSLCertificateInfo, error) {
	block, _ := pem.Decode(cer)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return x509Certificate(certificate), nil
}
