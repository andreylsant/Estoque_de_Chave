package token

import (
	"time"

	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	userEmail string `json:"name"`
	userID    string `json:"id"`
	jwt.RegisteredClaims
}

type ClaimsUser struct {
	Login login.Login
}

// Precisa de uma chave segreta
var chaveJWT = []byte("abcdefghijklnmopqrstuvwyz123456789")

func (c *ClaimsUser) CriarTokenJwt() (string, error) {
	//Adicionando tempo de expiração para meu tokem jwt
	tempo_Expiracao := time.Now().Add(24 * time.Hour)

	//Primeiro criar um claims
	claims := Claims{
		userEmail: c.Login.Email,
		userID:    c.Login.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tempo_Expiracao),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	//Criando token com a assinatura ES256
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(chaveJWT)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
