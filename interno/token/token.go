package token

import (
	"log/slog"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	userID    string `json:"id"`
	jwt.RegisteredClaims
}

type ServiceClaims struct{}

// Precisa de uma chave segreta
var chaveJWT = []byte("abcdefghijklnmopqrstuvwyz123456789")

func (s *ServiceClaims) GerarToken(userID string) (string, error) {
	slog.Info("Iniciando geração de token jwt")
	//Adicionando tempo de expiração para meu tokem jwt
	tempo_Expiracao := time.Now().Add(24 * time.Hour)

	//Primeiro criar um claims
	claims := Claims{
		userID:    userID,
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
