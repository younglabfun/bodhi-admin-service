package utils

import (
	"github.com/google/uuid"
	"github.com/thinkeridea/go-extend/exnet"
	"math/rand"
	"net/http"
)

var LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// CreateUuid create a version 4 UUID
func CreateUuid() string {
	uuidStr := uuid.Must(uuid.NewUUID())
	return uuidStr.String()
}

func GetRemoteIp(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}

func RandomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(b)
}

func ContainsStr(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
