package estoquekey

import (
	"fmt"
)

type Key struct {
	Id      string `json:"id"`
	Marca   string `json:"marca"`
	Modelo  int    `json:"modelo"`
	Unidade int    `json:"unidade"`
}

func NewKey(marca string, modelo, unidade int) (*Key, error) {
	key:= &Key{
		Id: "1",
		Marca:   marca,
		Modelo:  modelo,
		Unidade: unidade,
	}

	//chamando metodo error 
	err:= key.ValidarError()
	if err != nil{
		return nil, err
	}

	return key, nil
}

func (k *Key) ValidarError() error{
		if k.Marca == "" || k.Modelo == 0 {
		return  fmt.Errorf("[Error ao Criar marca]")
	} else if k.Unidade == 0 {
		return fmt.Errorf("[Error De Unidade]")
	}

	return nil
}