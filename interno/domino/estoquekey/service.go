package estoquekey

import "log/slog"

type Service struct {
	Repository Repository
}

func (s *Service) Create(chave *Chave) error {
	slog.Info("Integrando novas chaves a  tabela")
	newChave, err := NewChave(chave.Marca, chave.Modelo, chave.Quantidade)
	if err != nil {
		return err
	}

	slog.Info("Salvando novas chaves no banco de dados")
	err = s.Repository.Save(newChave)
	if err != nil {
		return err
	}

	return nil
}


func (s *Service) Get() ([]Chave, error) {
	chave, err:= s.Repository.Get()
	if err != nil{
		return nil, err
	}

	return chave, nil
}