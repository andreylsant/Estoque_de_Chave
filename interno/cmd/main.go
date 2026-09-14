package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andreylsant/estoque_de_chave/interno/controlle/handle"
	"github.com/andreylsant/estoque_de_chave/interno/db/database"
	"github.com/andreylsant/estoque_de_chave/interno/domino/estoquekey"
	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/andreylsant/estoque_de_chave/interno/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := &database.Salvarlogin{}
	token := &token.ServiceClaims{}
	service := login.LoginService{
		RepositoryLogin: db,
		Token:           token,
	}

	dbChave:= database.DBDeChave{}
	serviceChave:= estoquekey.Service{
		Repository: &dbChave,
	}
	handleChave:= handle.HandleChave{
		Service: serviceChave,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handle := handle.HandleLogin{
		LoginService: service,
	}
	router.POST("/criar-conta", handle.Post_Login)
	router.POST("/login", handle.LoginDeUsuario)
	router.POST("/produtos", handleChave.PostChave)
	router.GET("/produtos", handleChave.GetChave)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error na porta %s", err)
	}
}
