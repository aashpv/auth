package models

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("horovodova_help")

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
