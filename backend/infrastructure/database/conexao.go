package database

import (
	"fmt"
	"os"
	"strings"

	estoquechave "github.com/andreylsant/estoque_de_chave/backend/domain/estoque_chave"
	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConectaDB abre a conexão PostgreSQL e prepara as tabelas da aplicação.
// Em provedores como Neon, DATABASE_URL deve conter a URL completa com SSL.
func ConectaDB() (*gorm.DB, error) {
	dsn := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if dsn == "" {
		user := strings.TrimSpace(os.Getenv("POSTGRES_USER"))
		password := os.Getenv("POSTGRES_PASSWORD")
		databaseName := strings.TrimSpace(os.Getenv("POSTGRES_DB"))
		if user == "" || password == "" || databaseName == "" {
			return nil, fmt.Errorf("configure DATABASE_URL ou POSTGRES_USER, POSTGRES_PASSWORD e POSTGRES_DB")
		}

		sslMode := valorOuPadrao("DB_SSLMODE", "disable")
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			valorOuPadrao("DB_HOST", "localhost"),
			valorOuPadrao("DB_PORT", "5432"),
			user,
			password,
			databaseName,
			sslMode,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("conectar ao PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("obter conexão SQL: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("verificar conexão com PostgreSQL: %w", err)
	}

	if err := db.AutoMigrate(&loginuser.Login{}, &estoquechave.Chave{}); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("migrar tabelas: %w", err)
	}

	return db, nil
}

func valorOuPadrao(chave, padrao string) string {
	valor := strings.TrimSpace(os.Getenv(chave))
	if valor == "" {
		return padrao
	}
	return valor
}
