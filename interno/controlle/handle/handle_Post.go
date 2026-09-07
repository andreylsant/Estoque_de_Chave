package handle

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/gin-gonic/gin"
)

type HandleKey struct {
	LoginService login.LoginService
}

func (h *HandleKey) Post_Login(c *gin.Context) {
	fmt.Println("CHEGOU NO BACKEND!")

	var login login.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(500, gin.H{
			"erro": "Dados de login inválidos",
		})
		return
	}
	//
	slog.Info("inicindo service")
	service, err := h.LoginService.CreateLogin(&login)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	slog.Info("logando")
	slog.Info("logando no site.")
	c.JSON(http.StatusOK, gin.H{
		"token":    service,
		"mensagem": "Login recebido",
	})
}
