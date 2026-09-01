package login

import (
	"fmt"
	"log/slog"
)

type LoginService struct {
	RepositoryLogin RepositoryLogin
}

func (l *LoginService) CreateLogin(login *Login) error {
	//crio meu login 
	slog.Info("Login sendo criado!!")
	newLogin := NewLogin(login.Email, login.Senha)

	//Gero meu hash de senha 
	_, err := newLogin.HashSenha(newLogin.Senha)
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


func (l *LoginService) Logando(login *Login) (error) {
	//Vou precisa comparar ser meu login está correto 
	//Vou precisar comparar minha senha
	return nil
}