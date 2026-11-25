package handlechave

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/usecase/service"
	"github.com/gin-gonic/gin"
)

type Handle struct {
	Service service.Service
}

func (h *Handle) PostChave(c *gin.Context) {
	slog.Info("Iniciando meu postChave")
	var chave gormrepo.Chave

	if err := c.ShouldBindJSON(&chave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados do produto inválidos",
		})
		return
	}

	err := h.Service.Create(&chave)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao cadastrar produto",
		})
		return
	}

	c.JSON(201, gin.H{
		"mensagem": "Produto cadastrado com sucesso",
	})
}
