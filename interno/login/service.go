package login

import (
	"fmt"
	"log/slog"
)

type LoginService struct {
	RepositoryLogin RepositoryLogin
}

func (l *LoginService) CreateLogin(login *login) error {
	//crio meu login 
	slog.Info("Login sendo criado!!")
	newLogin := NewLogin(login.email, login.senha)

	//Gero meu hash de senha 
	_, err := newLogin.HashSenha(newLogin.senha)
	if err != nil{
		return err
	}

	//Salvo login no banco de dados 
	err = l.RepositoryLogin.SaveLogin(newLogin)
	if err != nil {
		return fmt.Errorf("[Error ao Salvar login]", err)
	}

	return nil
}

func (l *LoginService) GerarToken(){
	
}