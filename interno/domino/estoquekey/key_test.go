package estoquekey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	marca   = "stam"
	modelo  = "799"
	quantidade = 1000
)

func Test_CreateNewKey(t *testing.T) {
	assert := assert.New(t)
	newKey, _ := NewChave(marca, modelo,quantidade)

	assert.Equal(marca, newKey.Marca)
	assert.Equal(modelo, newKey.Modelo)
	assert.Equal(quantidade, newKey.Quantidade)
}

func Test_CreateNewKey_ErrorMarca(t *testing.T) {
	assert := assert.New(t)
	_, err := NewChave("", modelo, quantidade)

	assert.EqualError(err, "[Error ao Criar marca]")
}

func Test_CreateNewKey_ErrorModelo(t *testing.T) {
	assert := assert.New(t)
	_, err := NewChave(marca, "", quantidade)

	assert.EqualError(err, "[Error ao Criar marca]")
}

func Test_CreateNewKey_ErrorUnidade(t *testing.T) {
	assert := assert.New(t)
	_, err := NewChave(marca, modelo, 0)

	assert.EqualError(err, "[Error De Unidade]")
}
