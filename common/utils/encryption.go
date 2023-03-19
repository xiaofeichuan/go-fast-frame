package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type EncryptionUtil struct{}

var Encryption = new(EncryptionUtil)

// md5加密
func (*EncryptionUtil) Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
