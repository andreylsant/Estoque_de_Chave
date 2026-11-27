package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}

type ServiceClaims struct {
	secret []byte
}

// NewServiceClaims cria o gerador de tokens com um segredo fornecido pelo ambiente.
func NewServiceClaims(secret string) (*ServiceClaims, error) {
	secret = strings.TrimSpace(secret)
	if len(secret) < 32 {
		return nil, errors.New("JWT_SECRET precisa ter pelo menos 32 caracteres")
	}
	return &ServiceClaims{secret: []byte(secret)}, nil
}

func (s *ServiceClaims) GerarToken(userID string) (string, error) {
	if s == nil || len(s.secret) == 0 {
		return "", errors.New("gerador JWT não configurado")
	}
	if strings.TrimSpace(userID) == "" {
		return "", errors.New("identificador do usuário não informado")
	}

	now := time.Now()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("gerar token JWT: %w", err)
	}
	return tokenString, nil
}
