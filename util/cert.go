package util

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mundipagg/boleto-api/config"
)

func ListCert() error {

	list := []string{
		config.Get().CertBoletoPathCrt,
		config.Get().CertBoletoPathKey,
		config.Get().CertBoletoPathCa,
		config.Get().CertICP_PathPkey,
		config.Get().CertICP_PathChainCertificates,
	}

	var err error

	for _, v := range list {

		fmt.Println(v)
		err = copyCert(v)

		if err != nil {
			return err
		}

	}

	return err

}

func copyCert(c string) error {
	execPath, _ := os.Getwd()

	f := strings.Split(c, "/")

	fName := f[len(f)-1]

	srcFile, err := os.Open(execPath + "/boleto_orig/" + fName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(execPath + "/boleto_cert/" + fName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	err = destFile.Sync()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	err = os.Chmod(execPath+"/boleto_cert/"+fName, 0777)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}

	return err
}
