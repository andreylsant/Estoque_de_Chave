package handle

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/gin-gonic/gin"
)

type HandleKey struct {
	LoginService login.LoginService
}

func (h *HandleKey) Post_Login(c *gin.Context) {

	var login login.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(500, gin.H{
			"erro": "Dados de login inválidos",
		})
		return
	}
	//
	slog.Info("inicindo service")
	token, err := h.LoginService.CreateLogin(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	slog.Info("logando")
	slog.Info("logando no site.")
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"mensagem": "Login recebido",
	})
}
