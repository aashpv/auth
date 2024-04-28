// for middleware
package service

import (
	"errors"
	"fmt"
	"github.com/aashpv/auth/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		return claims.Email, nil
	}

	return "", errors.New("invalid auth token")
}
