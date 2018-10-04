package util

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	s "github.com/fullsailor/pkcs7"
	"github.com/mundipagg/boleto-api/config"
)

//Read privatekey and parse to PKCS#1
func parsePrivateKey() (crypto.PrivateKey, error) {

	pkeyBytes, err := ioutil.ReadFile(config.Get().CertICP_PathPkey)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pkeyBytes)
	if block == nil {
		return nil, errors.New("Key Not Found")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return rsa, nil
	default:
		return nil, fmt.Errorf("SSH: Unsupported key type %q", block.Type)
	}

}

///Read chainCertificates and adapter to x509.Certificate
func parseChainCertificates() (*x509.Certificate, error) {

	chainCertsBytes, err := ioutil.ReadFile(config.Get().CertICP_PathChainCertificates)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(chainCertsBytes)
	if block == nil {
		return nil, errors.New("Key Not Found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

// Read signedData and parse to *x509.Certificate
func parseSignedData(request string) (*s.SignedData, error) {

	sig, err := s.NewSignedData([]byte(request))

	return sig, err
}
