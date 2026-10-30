package handle

import (
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/domino/estoquekey"
	"github.com/gin-gonic/gin"
)

func (h *HandleChave) Add(c *gin.Context) {
	var chave estoquekey.Chave

	if err := c.ShouldBindJSON(&chave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados do produto inválidos",
		})
		return
	}

	h.Service.ServiceUpadete(&chave)
}

func (h *HandleChave) Delete_Chave(c *gin.Context){
	
}
