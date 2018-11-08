package service

import (
	"time"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/JiangInk/market_monitor/models"
	goJWT "github.com/dgrijalva/jwt-go"
)

// AuthService 认证相关
type AuthService struct{}

// GenerateToken 生成 Token
func (as AuthService) GenerateToken(user models.User) (string, error) {
	jwtInstance := jwt.NewJWT()
	nowTime := time.Now()
	expireTime := time.Duration(config.ServerConf.JWTExpire)
	claims := jwt.CustomClaims{
		ID: user.ID,
		UserName: user.UserName,
		Email: user.Email,
		StandardClaims: goJWT.StandardClaims {
			ExpiresAt: nowTime.Add(expireTime * time.Hour).Unix(),
			Issuer:    "monitor",
		},
	}
	return jwtInstance.CreateToken(claims)
	// todo set redis
}

// DestroyToken 销毁 Token
func (as AuthService) DestroyToken(token string) {
	return
}
