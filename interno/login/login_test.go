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

	newLogin, _:= NewLogin(email, senha)

	assert.Equal(newLogin.Email, email)
	assert.Equal(newLogin.Senha, senha)
}

func Test_ValidadeLogin(t *testing.T) {
	assert := assert.New(t)

	newLogin, _ := NewLogin(email, senha)
	err := newLogin.ValidadeLogin()

	assert.Equal(err, nil)
}

func Test_ValidadeLogin_ErrorLogin(t *testing.T) {
	assert := assert.New(t)

	newLogin, _:= NewLogin("", senha)
	err := newLogin.ValidadeLogin()

	assert.Equal(err, fmt.Errorf("error ao validar email e senha!"))
}

func Test_ValidadeLogin_Email(t *testing.T) {
	assert := assert.New(t)

	newLogin, _:= NewLogin(email, senha)
	err := newLogin.VerificarEmail()

	assert.Equal(err, fmt.Errorf("error ao criar email! Email invalido"))
}
