package utils

import (
	"go-fast-frame/constant"
	"go-fast-frame/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtil struct{}

var Jwt = new(JwtUtil)

type UserAuthClaims struct {
	UserId   uint              //用户Id
	UserName string            //用户账号
	NickName string            //用户昵称
	UserType constant.UserType //用户类型（0普通账号，1超级管理员）
	jwt.StandardClaims
}

// 签名
// var secretKey = []byte(global.AppConfig.Jwt.Signature)

// GenerateToken 生成token
func GenerateToken(claim UserAuthClaims, expireTime time.Time) (string, error) {

	claim.StandardClaims = jwt.StandardClaims{
		// 签名时间
		IssuedAt: time.Now().Unix(),
		// 过期时间
		ExpiresAt: expireTime.Unix(),
		// 签名颁发者
		Issuer: "go-fast-admin",
		// 签名主题
		Subject: "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	var secretKey = []byte(global.AppConfig.Jwt.Signature)
	return token.SignedString(secretKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	var secretKey = []byte(global.AppConfig.Jwt.Signature)
	// 解析 token string 拿到 token jwt.Token
	var claim UserAuthClaims
	return jwt.ParseWithClaims(tokenStr, &claim, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
