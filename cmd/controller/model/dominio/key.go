package dominio

import "errors"

type Key struct {
	Id      string `json:"id"`
	Marca   string `json:"marca"`
	Modelo  int    `json:"modelo"`
	Unidade int    `json:"unidade"`
}

func NewKey(marca string, modelo, unidade int) (*Key, error) {
	if marca == "" || modelo == 0 {
		return nil, errors.New("Error ao criar a chave")
	} else if unidade == 0 {
		return &Key{}, nil
	}

	return &Key{
		Id: "1",
		Marca:   marca,
		Modelo:  modelo,
		Unidade: unidade,
	}, nil
}
