package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
)

// MD5 md5 for string
func MD5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// UUID generate a uuid string
func UUID() string {
	id, _ := uuid.NewV4()
	return id.String()
}