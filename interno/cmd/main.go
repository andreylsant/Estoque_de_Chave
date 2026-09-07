package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andreylsant/estoque_de_chave/interno/controlle/handle"
	"github.com/andreylsant/estoque_de_chave/interno/db/database"
	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/andreylsant/estoque_de_chave/interno/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db:= &database.Salvarlogin{}
	token:= &token.ServiceClaims{}
	service:= login.LoginService{
		RepositoryLogin: db,
		Token: token,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handle := handle.HandleKey{
		LoginService: service,
	}
	router.POST("/login", handle.Post_Login)	

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error na porta %s", err)
	}
}
