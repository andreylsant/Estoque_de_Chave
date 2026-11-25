package database

import (
	"fmt"
	"os"

	estoquechave "github.com/andreylsant/estoque_de_chave/backend/domain/estoque_chave"
	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		valorOuPadrao("DB_HOST", "localhost"),
		valorOuPadrao("DB_PORT", "5432"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("conectar ao PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("obter conexão SQL: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("verificar conexão: %w", err)
	}

	if err := db.AutoMigrate(
		&loginuser.Login{},
		&estoquechave.Chave{},
	); err != nil {
		return nil, fmt.Errorf("migrar tabelas: %w", err)
	}

	return db, nil
}

func valorOuPadrao(chave, padrao string) string {
	valor := os.Getenv(chave)
	if valor == "" {
		return padrao
	}
	return valor
}
