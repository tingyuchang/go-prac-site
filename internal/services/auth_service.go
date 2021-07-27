package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-prac-site/internal/models"
	"time"
)

func CreateTokenString(username, password string)(string, error) {
	// 1 days
	expiredAt := time.Now().Add(24*time.Hour)
	payload := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			Issuer: "MattC",
		},
		Username: username,
		Password: password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte("token_password"))
}

func Check(tokenString string)(*models.Claims, error)  {
	tokenClaims,err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token*jwt.Token) (interface{}, error) {
		return []byte("token_password"), nil
	})

	if err != nil {
		return &models.Claims{}, err
	}

	if tokenClaims == nil {
		return &models.Claims{}, fmt.Errorf("tokenClaims is nil")
	}

	claims, ok := tokenClaims.Claims.(*models.Claims)

	if ok == false || tokenClaims.Valid == false {
		return &models.Claims{}, fmt.Errorf("tokenClaims is invalid")
	}

	// check claim: expiredAt
	if time.Now().Unix() > claims.ExpiresAt {
		return &models.Claims{}, fmt.Errorf("username is invalid or expired")
	}

	return claims, nil
}
