package controller

import (
	"dxcserver/common"
	"dxcserver/core"
	"dxcserver/model"
	"encoding/json"
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Config = common.LoadConfig()

func parseToken(token string) (*jwt.StandardClaims, error) {
	payload := strings.Split(token, ".") // 初步分割Token
	bytes, err := jwt.DecodeSegment(payload[1])
	if err != nil {
		return nil, err
	}
	// Parse Token
	var TokenDecode model.TokenDecode
	err = json.Unmarshal(bytes, &TokenDecode)
	if err != nil {
		return nil, err
	}
	// 刷新token
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(Config.JWT.Secret), nil
		})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			claim.Audience = TokenDecode.Aud
			return claim, nil
		}
	}

	return nil, err
}

func CreateToken(username, subject string) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  username,
		ExpiresAt: time.Now().Unix() + int64(60*60*24), // 24H时效
		IssuedAt:  time.Now().Unix(),
		Issuer:    "DaoxiangCity-Server",
		NotBefore: time.Now().Unix(),
		Subject:   subject,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(Config.JWT.Secret)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func RefreshToken(ctx *gin.Context) (string, string, error) {

	// 解析token获取username
	reqToken := ctx.DefaultQuery("Token", "")

	if reqToken == "" {
		return "", "", errors.New("token param is empty")
	}

	reqTokenClaims, err := parseToken(reqToken)

	if err != nil {
		return "", "", err
	}
	username := reqTokenClaims.Audience

	token, err := CreateToken(username, "TokenRefresh")
	if err != nil {
		return "", "", err
	}

	return username, token, nil
}

func Init() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.POST("/Config", core.ShowClientConfig)
	return engine
}
