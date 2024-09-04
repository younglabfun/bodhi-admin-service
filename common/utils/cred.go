package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func getMd5(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}

func GetMd5HexDigest(s string) string {
	b := getMd5([]byte(s))
	res := hex.EncodeToString(b)
	return res
}

func GetHashedPassword(password, salt string) string {
	res := GetMd5HexDigest(password)
	if salt != "" {
		res = GetMd5HexDigest(res + salt)
	}
	return res
}

func CheckPassword(plainPwd, hashedPwd, salt string) bool {
	return hashedPwd == GetHashedPassword(plainPwd, salt)
}
