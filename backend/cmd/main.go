package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	ginhandler "github.com/andreylsant/estoque_de_chave/backend/adapters/ginhandler"
	handlechave "github.com/andreylsant/estoque_de_chave/backend/adapters/ginhandler/handle_chave"
	handlelogin "github.com/andreylsant/estoque_de_chave/backend/adapters/ginhandler/handle_login"
	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/usecase/service"
	"github.com/andreylsant/estoque_de_chave/backend/infrastructure/database"
	"github.com/andreylsant/estoque_de_chave/backend/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// O arquivo .env é opcional: em hospedagem, as variáveis são configuradas no painel.
	if err := godotenv.Load(); err != nil {
		slog.Info("Arquivo .env não encontrado; usando variáveis do ambiente")
	}

	db, err := database.ConectaDB()
	if err != nil {
		log.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	jwtService, err := token.NewServiceClaims(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatalf("configuração JWT inválida: %v", err)
	}

	loginRepository := gormrepo.NewLoginRepository(db)
	loginService := service.LoginService{
		RepositoryLogin: loginRepository,
		Token:           jwtService,
	}

	chaveRepository := gormrepo.NewDBDeChave(db)
	chaveService := service.Service{Repository: chaveRepository}
	chaveHandler := handlechave.Handle{Service: chaveService}
	loginHandler := handlelogin.HandleLogin{LoginService: loginService}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/criar-conta", loginHandler.Post_Login)
	router.POST("/login", loginHandler.LoginDeUsuario)

	// O frontend já envia o JWT no cabeçalho Authorization após o login.
	produtos := router.Group("/produtos")
	produtos.Use(ginhandler.AuthRequired())
	produtos.POST("", chaveHandler.PostChave)
	produtos.PATCH("/:id/estoque", chaveHandler.AdicionarChave)
	produtos.GET("", chaveHandler.GetChave)

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("API iniciada na porta %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("falha ao iniciar servidor: %v", err)
	}
}

func allowedOrigins() []string {
	origins := strings.TrimSpace(os.Getenv("FRONTEND_URL"))
	if origins == "" {
		return []string{"http://localhost:5173"}
	}

	parts := strings.Split(origins, ",")
	result := make([]string, 0, len(parts))
	for _, origin := range parts {
		if origin = strings.TrimSpace(origin); origin != "" {
			result = append(result, origin)
		}
	}
	return result
}
