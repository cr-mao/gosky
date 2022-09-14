package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha1(data string) string {
	sha := sha1.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum(nil))
}

func Sha256(data string) string {
	sha := sha256.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum(nil))
}

func Md5(data string) string {
	md := md5.New()
	md.Write([]byte(data))
	return hex.EncodeToString(md.Sum(nil))
}
