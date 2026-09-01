package handle

import (
	"fmt"
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

	service, err := h.LoginService.CreateLogin(&login)

	if err != nil {
		fmt.Errorf("Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    service,
		"mensagem": "Login recebido",
	})
}
