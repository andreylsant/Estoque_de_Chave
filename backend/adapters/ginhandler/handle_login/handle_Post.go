package handlelogin

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/usecase/service"
	"github.com/gin-gonic/gin"
)

type HandleLogin struct {
	LoginService service.LoginService
}

var usuarioLogin gormrepo.Login

func (h *HandleLogin) Post_Login(c *gin.Context) {
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

func (h *HandleLogin) LoginDeUsuario(c *gin.Context){
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
