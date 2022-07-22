package tools

import (
	"github.com/dgrijalva/jwt-go"
)

type jwtParse[T any] struct {
	jwt.MapClaims
	Val T `json:"val"`
}

func JwtGenerate(sign string, c any) (string, error) {
	tok := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"val": c,
	})
	return tok.SignedString([]byte(sign))
}

func JwtParse[T any](sign string, token string) (T, error) {

	tokenInfo := jwtParse[T]{}
	_, _ = jwt.ParseWithClaims(token, &tokenInfo, func(token *jwt.Token) (any, error) {
		return []byte(sign), nil
	})

	return tokenInfo.Val, nil
}
