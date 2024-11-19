package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

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

func RandomStr(n int) string {
	b := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		b = append(b, LETTERS[r.Intn(len(LETTERS))])
	}
	return string(b)
}
