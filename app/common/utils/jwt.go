package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/king-tu/mallweb/app/common"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/models"
	"go.uber.org/zap"
	"time"
)

func GenerateToken(user *models.User, expiresTime int, secrect string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * time.Duration(expiresTime))

	claims := jwt.StandardClaims{
		Audience:  user.Name,           // 受众
		ExpiresAt: expireTime.Unix(),   // 失效时间
		Id:        string(user.Id),     // 编号
		IssuedAt:  nowTime.Unix(),      // 签发时间
		Issuer:    common.PROJECT_NAME, // 签发人
		NotBefore: nowTime.Unix(),      // 生效时间
		Subject:   "login",             // 主题
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secrect))
	if err != nil {
		global.Logger.Bg().Error(false, "Fail to generating access token", zap.Error(err))
		return "", err
	}

	return token, nil
}

func ParseToken(token, secret string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
