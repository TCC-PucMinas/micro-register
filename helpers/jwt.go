package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
	hmacSecret      = []byte(os.Getenv("HMAC_SECRET"))
)

var TimeToken = 30
var timeRefresh = 1000

type Claims struct {
	Username string `json:"id"`
	Sub      int    `json:"sub"`
	jwt.StandardClaims
}

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Claims) GenerateClains(id string, sub int, minutes int64) {
	c.Username = id
	c.Sub = sub
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
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
	claimsToken := &Claims{}
	claimsRefresh := &Claims{}
	jwtToken := JwtTokens{}
	claimsToken.GenerateClains(id, 0, int64(TimeToken))
	accessToken, err := tokenHash(*claimsToken)

	if err != nil {
		return jwtToken, err
	}

	claimsRefresh.GenerateClains(id, 1, int64(timeRefresh))
	refreshToken, er := tokenHash(*claimsRefresh)

	if er != nil {
		return jwtToken, er
	}

	jwtToken.AccessToken = accessToken
	jwtToken.RefreshToken = refreshToken

	return jwtToken, nil
}

func CheckJwt(jwtUser string, sub int) (bool, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtUser, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return false, ErrExpiredToken
		}

		return false, ErrInvalidToken
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid && sub == claims.Sub {
		return true, nil
	}

	return false, err
}

func ExtractJwt(jwtUser string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtUser, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if err != nil {
		return claims, err
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return claims, errors.New("Error extract jwt")
}
