package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandleChave) GetChave(c *gin.Context) {
	chaves, err := h.Service.Get()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar produtos",
		})
		return
	}

	c.JSON(http.StatusOK, chaves)
}
