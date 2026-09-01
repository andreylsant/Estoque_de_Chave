package login

import (
	"fmt"
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/interno/token"
)

type LoginService struct {
	RepositoryLogin RepositoryLogin
}

func (l *LoginService) CreateLogin(login *Login) (string, error) {
	claims := token.ClaimsUser{}

	//crio um novo login
	slog.Info("Login sendo criado!!")
	newLogin := NewLogin(login.Email, login.Senha)

	//Gero meu hash de senha
	_, err := newLogin.HashSenha(newLogin.Senha)
	if err != nil {
		return "", err
	}

	//Salvo login no banco de dados
	err = l.RepositoryLogin.SaveLogin(newLogin)
	if err != nil {
		return "", fmt.Errorf("[Error ao Salvar login]", err)
	}

	token, err := claims.CriarTokenJwt()
	if err != nil {
		return "", fmt.Errorf("[Error na criação do token]", err)
	}
	return token, nil
}

func (l *LoginService) Logando(login *Login) error {
	//Vou precisa comparar ser meu login está correto
	//Vou precisar comparar minha senha
	return nil
}
