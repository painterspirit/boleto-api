package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

const _key = "2131231231231234"

//Encrypt encripta texto baseado na documentação do GO
func Encrypt(s string) string {
	key := []byte(_key)
	plaintext := []byte(s)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

//Decrypt decripta string encriptada
func Decrypt(s string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(s)

	block, err := aes.NewCipher([]byte(_key))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

//Base64 converte um string para base64
func Base64(s string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	return sEnc
}

//Base64Decode converte uma string base64 para uma string normal
func Base64Decode(s string) string {
	sDec, _ := base64.StdEncoding.DecodeString(s)
	return string(sDec)
}

//Sha256 converts string to hash sha256. Encode default: Base64
func Sha256(s string, e string) string {
	var sEnc string

	h := sha256.New()
	h.Write([]byte(s))

	switch e {
	case "hex":
		sEnc = hex.EncodeToString(h.Sum(nil))
	default:
		sEnc = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}
	return sEnc
}
