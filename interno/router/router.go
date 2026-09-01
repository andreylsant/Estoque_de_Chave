package router

import (
	"fmt"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/controlle/handle"
	"github.com/gin-gonic/gin"
)

//criando minhas router, mas não lembro se a func
func MyRouter(r *gin.Engine){
	handle:= handle.HandleKey{}
	if err:= http.ListenAndServe("8080", nil); err != nil{
		fmt.Println("Error na porta  8080")
	}

	r.POST("/login", handle.Post_Login)
}