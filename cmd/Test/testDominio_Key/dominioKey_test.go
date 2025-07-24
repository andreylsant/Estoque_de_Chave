package testdominiokey

import (
	"fmt"
	"testing"

	"github.com/andreylsant/estoque_de_chave/cmd/controller/model/dominio"
)

var (
	marca = "stam"
	modelo = 799
	unidade = 19
)

func Test_NewKey(t *testing.T) {
	newKey, _ := dominio.NewKey(marca, modelo, unidade)

	if newKey.Marca == "3f" {
		fmt.Println("Marca igual  a marca")
	} else {
		fmt.Println("Não são iguais")
	}
}

func Test_NewUnidade_Error(t *testing.T) {
	//_, err := dominio.NewKey(marca, modelo, 0)
}


