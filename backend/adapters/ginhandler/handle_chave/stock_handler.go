package handlechave

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/gin-gonic/gin"
)

func (h *Handle) AdicionarChave(c *gin.Context) {
	slog.Info("Recebendo dados de entrada!")
	var chave gormrepo.Chave

	if err := c.ShouldBindJSON(&chave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados do produto inválidos",
		})
		return
	}

	h.Service.ServiceUpadete(&chave)
}

func (h *Handle) DiminiurChave(c *gin.Context) {

}
