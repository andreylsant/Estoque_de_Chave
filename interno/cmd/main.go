package main

import (
	"fmt"
	"net/http"

	"github.com/andreylsant/estoque_de_chave/interno/controlle/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()
	handle := handle.HandleKey{}
	err := http.ListenAndServe(":3000", nil)
	if  err != nil {
		fmt.Println("Error na porta %s", err) 
	}

	router.POST("/login", handle.Post_Login)
}
