package database

import (
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

func (db *DBDeChave) GetKeyID(id string) *estoquekey.Chave {
	return nil
}

func (db *DBDeChave) DeleteID(chave *estoquekey.Chave) (string, error) {
	return "", nil
}
