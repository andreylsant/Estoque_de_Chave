package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	handlechave "github.com/andreylsant/estoque_de_chave/backend/adapters/ginhandler/handle_chave"
	handlelogin "github.com/andreylsant/estoque_de_chave/backend/adapters/ginhandler/handle_login"
	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/usecase/service"
	"github.com/andreylsant/estoque_de_chave/backend/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load("../../.env"); err != nil {
		slog.Error("Não foi possível carregar o arquivo .env", "erro", err)
		os.Exit(1)
	}

	db := &gormrepo.Salvarlogin{}
	token := &token.ServiceClaims{}
	service_l := service.LoginService{
		RepositoryLogin: db,
		Token:           token,
	}

	dbChave, err := gormrepo.NewDBDeChave()
	if err != nil {
		slog.Error("Erro ao conectar ao banco", "erro", err)
		os.Exit(1)
	}
	//dbChave:= database.DBDeChave{}
	serviceChave := service.Service{
		Repository: dbChave,
	}
	handleChave := handlechave.Handle{
		Service: serviceChave,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handle := handlelogin.HandleLogin{
		LoginService: service_l,
	}
	router.POST("/criar-conta", handle.Post_Login)
	router.POST("/login", handle.LoginDeUsuario)
	router.POST("/produtos", handleChave.PostChave)
	router.PATCH("/produtos/:id/entrada", handleChave.AdicionarChave)
	router.GET("/produtos", handleChave.GetChave)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error na porta %s", err)
	}
}
