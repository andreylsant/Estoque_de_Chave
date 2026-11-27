package service

import (
	"errors"
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/ports"
	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
)

type LoginService struct {
	RepositoryLogin ports.LoginRepository
	Token           TokenGerador
}

type TokenGerador interface {
	GerarToken(userID string) (string, error)
}

func (l *LoginService) CreateLogin(login *gormrepo.Login) (string, error) {

	//crio um novo login
	slog.Info("Login sendo criado!!")
	newLogin, err := loginuser.NewLogin(login.Email, login.Senha)
	if err != nil {
		return "", err
	}

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

	return newLogin.Email, nil
}

func (l *LoginService) ValidarLogin(login *gormrepo.Login) (string, error) {
	slog.Info("iniciando Validação de login")
	//Vou precisa comparar ser meu login está correto
	//vou receber o meu login. então  vou pegar o  email que foi enviado e verificar se ja teenho esse email
	//Vou precisar comparar minha senha

	/*
		↓
		Valida formato
		↓
		Valida domínio
		↓
		Busca email no banco
		↓
		Email existe?
		↓
		Compara senha com o hash
		↓
		Senha correta?
		↓
		Gera JWT
		↓
		Retorna token
	*/

	newLogin := login

	slog.Info("verificando se emial está cadastrado no banco")
	usuario, err := l.RepositoryLogin.VerificaEmail(newLogin.Email)
	if err != nil {
		return "", err
	}

	err = loginuser.CompararSenha(newLogin.Senha, usuario.Senha)
	if err != nil {
		return "", err
	}

	//Aqui estou passando meu token para o usuario
	slog.Info("geranado token")
	token, err := l.Token.GerarToken(usuario.Id)
	if err != nil {
		return "", errors.New("error ao criar token")
	}
	slog.Info("retornando token ao cliente")

	return token, nil
}
