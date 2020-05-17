package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/king-tu/mallweb/app/models"
	"github.com/king-tu/mallweb/common"
	"go.uber.org/zap"
	"time"
)

func GenerateToken(user *models.User, expiresTime time.Duration, secrect string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expiresTime)

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
	zap.L().Sugar().Debugf("The access token is %s", token)
	if err != nil {
		zap.L().Sugar().Errorf("The error during generating access token is %v", err)
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
