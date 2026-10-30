package database

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/interno/domino/estoquekey"
	"gorm.io/gorm"
)

type DBDeChave struct {
	DB *gorm.DB
}

func NewDBDeChave() (*DBDeChave, error) {
	db, err := ConectaDB()
	if err != nil {
		return nil, err
	}

	return &DBDeChave{
		DB: db,
	}, nil
}

func (db *DBDeChave) Save(chave *estoquekey.Chave) error {
	slog.Info("iniciando banco de dados")
	return db.DB.Create(chave).Error
}

func (db *DBDeChave) Upadete(chave *estoquekey.Chave) error {
	slog.Info("iniciando banco de dados")
	db.DB.Save(chave)
	return nil
}

func (db *DBDeChave) Get() ([]estoquekey.Chave, error) {
	var listaDeChave []estoquekey.Chave

	err := db.DB.Find(&listaDeChave).Error
	if err != nil {
		return nil, err
	}

	return listaDeChave, nil
}

func (db *DBDeChave) GetKeyID(id string) (*estoquekey.Chave, error) {
	var chave estoquekey.Chave

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
	result := db.DB.Delete(&estoquekey.Chave{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("excluir chave: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
