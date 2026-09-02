package database

import "github.com/andreylsant/estoque_de_chave/interno/login"

type Salvarlogin struct {
	login []login.Login
}

func (s *Salvarlogin) SaveLogin(login *login.Login) error{
	s.login = append(s.login, *login)
	return nil
}
