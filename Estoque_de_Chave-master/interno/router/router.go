package router

import (
	"github.com/andreylsant/estoque_de_chave/interno/controlle/handle"
	"github.com/gin-gonic/gin"
)

//criando minhas router, mas não lembro se a func 
func MyRouter(r *gin.Engine){
	handle:= handle.HandleKey{}

	r.POST("/login", handle.Login)
}