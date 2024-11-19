package utils

import (
	"github.com/google/uuid"
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

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
