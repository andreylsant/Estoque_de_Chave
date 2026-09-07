package login

import (
	"errors"
	"fmt"
	"log/slog"
)

type LoginService struct {
	RepositoryLogin RepositoryLogin
	Token TokenGerador
}

type TokenGerador interface {
    GerarToken(userID string) (string, error)
}	

func (l *LoginService) CreateLogin(login *Login) (string, error) {

	//crio um novo login
	slog.Info("Login sendo criado!!")
	newLogin := NewLogin(login.Email, login.Senha)

	//Gero meu hash de senha
	slog.Info("gerando hesh")
	hash, err := newLogin.HashSenha(newLogin.Senha)
	if err != nil {
		return "", err
	}
	newLogin.Senha = hash

	//Salvo login no banco de dados
	slog.Info("salvando no banco de dados")
	err = l.RepositoryLogin.SaveLogin(newLogin)
	if err != nil {
		return "", errors.New("[Error ao Salvar login]")
	}

	//Aqui estou passando meu token para o usuario 
	slog.Info("geranado token")
	token, err := l.Token.GerarToken(login.Id)
	if err != nil {
		return "", fmt.Errorf("[Error na criação do token]", err)
	}

	slog.Info("retornando token ao cliente")
	return token, nil
}

func (l *LoginService) Logando(login *Login) error {
	//Vou precisa comparar ser meu login está correto
	//Vou precisar comparar minha senha
	return nil
}
