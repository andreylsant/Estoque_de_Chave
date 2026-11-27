package gormrepo

import (
	"errors"
	"fmt"

	loginuser "github.com/andreylsant/estoque_de_chave/backend/domain/login_user"
	"gorm.io/gorm"
)

// LoginRepository implementa a porta de persistência de usuários usando GORM.
type LoginRepository struct {
	DB *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{DB: db}
}

func (r *LoginRepository) SaveLogin(login *loginuser.Login) error {
	if r == nil || r.DB == nil {
		return errors.New("repositório de login sem conexão com o banco")
	}
	if err := r.DB.Create(login).Error; err != nil {
		return fmt.Errorf("salvar usuário: %w", err)
	}
	return nil
}

func (r *LoginRepository) VerificaEmail(email string) (*loginuser.Login, error) {
	if r == nil || r.DB == nil {
		return nil, errors.New("repositório de login sem conexão com o banco")
	}

	var usuario loginuser.Login
	if err := r.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, fmt.Errorf("buscar usuário por email: %w", err)
	}
	return &usuario, nil
}
