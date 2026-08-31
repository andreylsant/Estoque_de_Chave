package login

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type login struct {
	id    int  
	email string 
	senha string 
}

func NewLogin(email, senha string) *login {
	l:= &login{
		id:    1,
		email: email,
		senha: senha,
	}

	l.ValidadeLogin()

	return l 
}

func (l *login) ValidadeLogin() error{
	slog.Info("Iniciando validação!!")

	if l.email == "" || l.senha == ""{
		return fmt.Errorf("error ao validar email e senha!")
	}

	return nil
}

func (l *login) HashSenha(senha string) (string, error){
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(senha),
		bcrypt.DefaultCost,
	)
	if err != nil{
		return  "", fmt.Errorf("Error ao gerar hash", err)
	}

	return string(hash), nil
}
