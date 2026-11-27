package estoquechave

import (
	"fmt"
	"time"

	"github.com/rs/xid"
)

type Chave struct {
	Id         string `gorm:"primaryKey;size:32"`
	Marca      string `gorm:"not null"`
	Modelo     string `gorm:"not null"`
	Quantidade int    `gorm:"not null;default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewChave(marca, modelo string, quantidade int) (*Chave, error) {
	chave := &Chave{
		Id:         xid.New().String(),
		Marca:      marca,
		Modelo:     modelo,
		Quantidade: quantidade,
		CreatedAt:  time.Now(),
	}

	//chamando metodo error
	err := chave.ValidarError()
	if err != nil {
		return nil, err
	}

	return chave, nil
}

func (k *Chave) ValidarError() error {
	if k.Marca == "" || k.Modelo == "" {
		return fmt.Errorf("[Error ao Criar marca]")
	} else if k.Quantidade == 0 {
		return fmt.Errorf("[Error De Unidade]")
	}

	return nil
}
