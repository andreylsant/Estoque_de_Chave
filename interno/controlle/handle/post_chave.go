package handle

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/domino/estoquekey"
	"github.com/gin-gonic/gin"
)

type HandleChave struct {
	Service estoquekey.Service
}

func (h *HandleChave) PostChave(c *gin.Context) {
	slog.Info("Iniciando meu postChave")
	var chave estoquekey.Chave

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

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Produto cadastrado com sucesso",
	})
}
