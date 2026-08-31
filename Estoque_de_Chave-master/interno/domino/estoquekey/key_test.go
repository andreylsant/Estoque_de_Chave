package estoquekey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	marca   = "stam"
	modelo  = 799
	unidade = 1000
)

func Test_CreateNewKey(t *testing.T) {
	assert := assert.New(t)
	newKey, _ := NewKey(marca, modelo, unidade)

	assert.Equal(marca, newKey.Marca)
	assert.Equal(modelo, newKey.Modelo)
	assert.Equal(unidade, newKey.Unidade)
}

func Test_CreateNewKey_ErrorMarca(t *testing.T) {
	assert := assert.New(t)
	_, err := NewKey("", modelo, unidade)

	assert.EqualError(err, "[Error ao Criar marca]")
}

func Test_CreateNewKey_ErrorModelo(t *testing.T) {
	assert := assert.New(t)
	_, err := NewKey(marca, 0, unidade)

	assert.EqualError(err, "[Error ao Criar marca]")
}

func Test_CreateNewKey_ErrorUnidade(t *testing.T) {
	assert := assert.New(t)
	_, err := NewKey(marca, modelo, 0)

	assert.EqualError(err, "[Error De Unidade]")
}
