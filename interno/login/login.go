package login

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Id    string
	Email string
	Senha string
}

func NewLogin(email, senha string) *Login {
	l := &Login{
		Id:    "1",
		Email: email,
		Senha: senha,
	}

	l.ValidadeLogin()

	return l
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
