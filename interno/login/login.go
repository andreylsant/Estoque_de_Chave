package login

import (
	"errors"
	"fmt"
	"log/slog"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var chipcode = "@chipcode.com"

type Login struct {
	Id    string `json:"id"`
	Email string `josn:"email"`
	Senha string `json:"senha"`
}

func NewLogin(email, senha string) (*Login, error) {
	slog.Info("criando login func newlogin()")
	l := &Login{
		Id:    "1",
		Email: email,
		Senha: senha,
	}

	err := l.ValidadeLogin()
	if err != nil {
		return nil, err
	}

	err= l.VerificarEmail()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Login) ValidadeLogin() error {
	slog.Info("Iniciando validação!!")

	if l.Email == "" || l.Senha == "" {
		return fmt.Errorf("error ao validar email e senha!")
	}

	return nil
}

func (l *Login) HashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(senha),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("Error ao gerar hash", err)
	}

	return string(hash), nil
}

func (l *Login) VerificarEmail() error {
	slog.Info("validado email...")

	_, err := mail.ParseAddress(l.Email)
	if err != nil {
		return errors.New("email inválido")
	}

	if !strings.HasSuffix(l.Email, chipcode) {
		return errors.New("email inválido: utilize um email @chipcode.com")
	}

	/*if !strings.Contains(l.Email, chipcode) {
		return errors.New("error ao criar email! Email invalido")
	}*/

	return nil
}
