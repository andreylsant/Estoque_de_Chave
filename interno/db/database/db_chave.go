package database

import (
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/interno/domino/estoquekey"
)

type DBDeChave struct {
	ListaDeChave []estoquekey.Chave
}

func (db *DBDeChave) Save(chave *estoquekey.Chave) error {
	slog.Info("iniciando banco de dados")
	db.ListaDeChave = append(db.ListaDeChave, *chave)
	return nil
}

func (db *DBDeChave) Upadete(chave *estoquekey.Chave) error {
	slog.Info("iniciando banco de dados")
	db.ListaDeChave = append(db.ListaDeChave, *chave)
	return nil
}

func (db *DBDeChave) Get() ([]estoquekey.Chave, error) {
	listaDeChave:= db.ListaDeChave
	return listaDeChave, nil
}

func (db *DBDeChave) GetKeyID(id string) *estoquekey.Chave {
	return nil
}
