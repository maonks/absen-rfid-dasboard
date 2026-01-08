package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

var SECRET_KEY []byte

func InitSecret() {
	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
}

func CekHMAC(body []byte, sign string) bool {

	InitSecret()
	m := hmac.New(sha256.New, SECRET_KEY)

	m.Write(body)
	return hmac.Equal([]byte(hex.EncodeToString(m.Sum(nil))), []byte(sign))

}
