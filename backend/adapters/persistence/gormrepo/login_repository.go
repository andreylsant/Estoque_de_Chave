package gormrepo

import (
	"errors"
	"log/slog"

	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
)

type Salvarlogin struct {
	Login []loginuser.Login
}

func (s *Salvarlogin) SaveLogin(login *loginuser.Login) error {
	slog.Info("iniciando banco de dados")
	s.Login = append(s.Login, *login)
	return nil
}

func (s *Salvarlogin) VerificaEmail(email string) (*loginuser.Login, error) {
	for _, use := range s.Login {
		if use.Email == email {
			return &use, nil
		}
	}
	return nil, errors.New("usuário não encontrado")
}
