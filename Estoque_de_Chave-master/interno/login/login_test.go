package login

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	senha = "senha"
	email = "121345656"
)

func Test_NewLogin(t *testing.T) {
	assert := assert.New(t)

	newLogin := NewLogin(email, senha)

	assert.Equal(newLogin.email, email)
	assert.Equal(newLogin.senha, senha)
}

func Test_ValidadeLogin(t *testing.T) {
	assert := assert.New(t)

	newLogin := NewLogin(email, senha)
	err := newLogin.ValidadeLogin()

	assert.Equal(err, nil)
}

func Test_ValidadeLogin_ErrorLogin(t *testing.T) {
	assert := assert.New(t)

	newLogin := NewLogin("", senha)
	err := newLogin.ValidadeLogin()

	assert.Equal(err, fmt.Errorf("error ao validar email e senha!"))
}
