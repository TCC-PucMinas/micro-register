package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSecret = []byte(os.Getenv("hmac_secret")) // []byte("e6428fcb1ea69f53bec5a2b4b817937b75db71d23207c6a6fb4b5cd2c9b9b43f")

var timeToken = 60
var timeRefresh = 1440

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Claims) GenerateClains(id string, minutes int64) {
	c = &Claims{
		Username: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
		},
	}

}

func tokenHash(claims Claims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString(hmacSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJwt(id string) (JwtTokens, error) {
	claims := &Claims{}
	jwtToken := JwtTokens{}
	claims.GenerateClains(id, int64(timeToken))
	accessToken, err := tokenHash(*claims)

	if err != nil {
		return jwtToken, err
	}

	claims.GenerateClains(id, int64(timeRefresh))
	refreshToken, er := tokenHash(*claims)

	if er != nil {
		return jwtToken, er
	}

	jwtToken.AccessToken = accessToken
	jwtToken.RefreshToken = refreshToken

	return jwtToken, nil

}

func CheckJwt(jwtUser string) (bool, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtUser, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, nil
	}

	return false, err
}
