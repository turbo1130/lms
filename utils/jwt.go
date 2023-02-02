package utils

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func MakeToken(user model.User) (tokenString string, err error) {
	ep, err := ParseTimeDur(global.LMS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return "", err
	}
	claim := request.CustomClaims{
		BaseClaims: request.BaseClaims{
			ID:       user.ID,
			Username: user.Username,
			NickName: user.NickName,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间，根据配置确定
			IssuedAt:  jwt.NewNumericDate(time.Now()),         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),         // 生效时间
			Issuer:    global.LMS_CONFIG.JWT.Issuer,           // 签发者
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)                       // 使用HS256算法
	tokenString, err = token.SignedString([]byte(global.LMS_CONFIG.JWT.SigningKey)) // 注意私钥只能用byte数组，string类型jwt无法解析
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(global.LMS_CONFIG.JWT.SigningKey), nil
	}
}

func ParseToken(t string) (*request.CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(t, &request.CustomClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("非法token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token到期，已失效")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token尚未激活")
			} else {
				return nil, errors.New("无法处理该token")
			}
		}
	}
	if claims, ok := parsedToken.Claims.(*request.CustomClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errors.New("无法处理该token")
}
