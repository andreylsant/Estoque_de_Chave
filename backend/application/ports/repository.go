package ports

import (
	estoquechave "github.com/andreylsant/estoque_de_chave/backend/domain/estoque_chave"
	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
)

type ChaveRepository interface {
	Save(*estoquechave.Chave) error
	Upadete(*estoquechave.Chave) error
	Get() ([]estoquechave.Chave, error)
	GetKeyID(id string) (*estoquechave.Chave, error)
	DeleteID(id string) error
}

type LoginRepository interface {
	SaveLogin(*loginuser.Login) error
	VerificaEmail(email string) (*loginuser.Login, error)
}

type PasswordHasher interface {
	Hash(string) (string, error)
	Compare(hash, plain string) error
}
type TokenGenerator interface {
	Generate(userID string) (string, error)
}
