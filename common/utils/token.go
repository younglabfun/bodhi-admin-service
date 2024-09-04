package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtAuth struct {
	Secret  string `json:"secret"`
	Expired int64  `json:"expired"`
}

type UserData struct {
	UserUuid string `json:"userUuid"`
}
type jwtToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

func BuildToken(jwtAuth JwtAuth, user UserData) *jwtToken {
	type payloads map[string]interface{}
	p := make(payloads)
	p["UserUuid"] = user.UserUuid

	jwtToken, err := BuildAccessToken(jwtAuth, p)
	if err != nil {
		return nil
	}
	refreshToken := BuildRefreshToken()
	jwtToken.RefreshToken = refreshToken
	return jwtToken
}

func BuildAccessToken(jwtAuth JwtAuth, payloads map[string]interface{}) (*jwtToken, error) {
	now := time.Now().Unix()
	accessToken, err := GenToken(now, jwtAuth.Secret, payloads, jwtAuth.Expired)
	if err != nil {
		return nil, err
	}

	return &jwtToken{
		AccessToken:  accessToken,
		AccessExpire: now + jwtAuth.Expired,
		RefreshAfter: now + jwtAuth.Expired/2,
	}, nil
}

func GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func BuildRefreshToken() string {
	now := time.Now().Unix()
	randomStr := RandomStr(16)
	md5Str := md5.Sum([]byte(randomStr + string(now)))
	refreshToken := hex.EncodeToString(md5Str[:])
	return refreshToken
}
