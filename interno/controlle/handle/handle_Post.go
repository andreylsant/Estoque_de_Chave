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

var usuarioLogin login.Login

func (h *HandleKey) Post_Login(c *gin.Context) {
	slog.Info("Criando usuario")
	if err := c.ShouldBindJSON(&usuarioLogin); err != nil {
		c.JSON(500, gin.H{
			"erro": "Dados de login inválidos",
		})
		return
	}
	//
	slog.Info("inicindo service")
	email, err := h.LoginService.CreateLogin(&usuarioLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":    email,
		"mensagem": "Email valido",
	})


}

func (h *HandleKey) LoginDeUsuario(c *gin.Context){
	slog.Info("iniciando loge")
	if err := c.ShouldBindJSON(&usuarioLogin); err != nil {
		c.JSON(500, gin.H{
			"erro": "Dados de login inválidos",
		})
		return
	}

	token, err:= h.LoginService.ValidarLogin(&usuarioLogin)
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
