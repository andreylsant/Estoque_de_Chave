package gormrepo

import (
	"errors"
	"fmt"
	"log/slog"

	estoquechave "github.com/andreylsant/estoque_de_chave/backend/domain/estoque_chave"
	"gorm.io/gorm"
)

type DBDeChave struct {
	DB *gorm.DB
}

func NewDBDeChave(db *gorm.DB) *DBDeChave {
	return &DBDeChave{DB: db}
}

func (db *DBDeChave) Save(chave *estoquechave.Chave) error {
	slog.Info("iniciando banco de dados")
	return db.DB.Create(chave).Error
}

func (db *DBDeChave) Upadete(chave *estoquechave.Chave) error {
	slog.Info("iniciando banco de dados")
	return db.DB.Save(chave).Error
}

func (db *DBDeChave) Get() ([]estoquechave.Chave, error) {
	var listaDeChave []estoquechave.Chave

	err := db.DB.Find(&listaDeChave).Error
	if err != nil {
		return nil, err
	}

	return listaDeChave, nil
}

func (db *DBDeChave) GetKeyID(id string) (*estoquechave.Chave, error) {
	var chave estoquechave.Chave

	result := db.DB.First(&chave, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("buscar chave por ID: %w", result.Error)
	}

	return &chave, nil
}

func (db *DBDeChave) DeleteID(id string) error {
	result := db.DB.Delete(&estoquechave.Chave{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("excluir chave: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
