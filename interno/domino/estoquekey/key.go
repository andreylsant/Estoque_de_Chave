package estoquekey

import (
	"fmt"
)

type Chave struct {
	Id      string `json:"id"`
	Marca   string `json:"marca"`
	Modelo  string    `json:"modelo"`
	Quantidade int   `json:"quantidade"`
}

func NewChave(marca, modelo string, quantidade int) (*Chave, error) {
	chave:= &Chave{
		Id: "1",
		Marca:   marca,
		Modelo:  modelo,
		Quantidade: quantidade,
	}

	//chamando metodo error 
	err:= chave.ValidarError()
	if err != nil{
		return nil, err
	}

	return chave, nil
}

func (k *Chave) ValidarError() error{
		if k.Marca == "" || k.Modelo == "" {
		return  fmt.Errorf("[Error ao Criar marca]")
	} else if k.Quantidade == 0{
		return fmt.Errorf("[Error De Unidade]")
	}

	return nil
}