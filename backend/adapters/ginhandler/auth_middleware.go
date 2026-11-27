package ginhandler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token não informado"})
			return
		}
		secret := os.Getenv("JWT_SECRET")
		if len(secret) < 32 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"erro": "JWT_SECRET não configurado corretamente"})
			return
		}
		claims := &tokenClaims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid || claims.UserID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido ou expirado"})
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
