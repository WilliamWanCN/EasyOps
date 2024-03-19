package common

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("easyops is very good.")
var TokenExpired = time.Hour * 12

type MyClaim struct {
	Username string
	jwt.StandardClaims
}

func GenToken(ctx *gin.Context, username string) (string, error) {
	myClaim := MyClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(TokenExpired),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "easyops",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 422,
			"msg":  err.Error(),
		})
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(ctx *gin.Context, tokenStr string) (*jwt.Token, *MyClaim, error) {
	myClaim := &MyClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, myClaim, func(t *jwt.Token) (interface{}, error) { return JwtKey, nil })
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 422,
			"msg":  err.Error(),
		})
		return nil, nil, err
	}
	return token, myClaim, nil
}

func EncryptPwd(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
func DecryptPwd(encryptPwd string) string {
	decoded, err := base64.StdEncoding.DecodeString(encryptPwd)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(decoded)
}
