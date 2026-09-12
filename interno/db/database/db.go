package database

import (
	"errors"
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/interno/login"
)

type Salvarlogin struct {
	login []login.Login
}

func (s *Salvarlogin) SaveLogin(login *login.Login) error {
	slog.Info("iniciando banco de dados")
	s.login = append(s.login, *login)
	return nil
}

func (s *Salvarlogin) VerificaEmail(email string) (*login.Login, error) {
	for _, use := range s.login {
		if use.Email == email {
			return &use, nil
		}
	}
	return nil, errors.New("usuário não encontrado")
}
