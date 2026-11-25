package service

import (
	"fmt"
	"log/slog"

	"github.com/andreylsant/estoque_de_chave/backend/adapters/persistence/gormrepo"
	"github.com/andreylsant/estoque_de_chave/backend/application/ports"
	estoquechave "github.com/andreylsant/estoque_de_chave/backend/domain/estoque_chave"
)

type Service struct {
	Repository ports.ChaveRepository
}

func (s *Service) Create(chave *gormrepo.Chave) error {
	slog.Info("Integrando novas chaves a  tabela")
	newChave, err := estoquechave.NewChave(chave.Marca, chave.Modelo, chave.Quantidade)
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

func (s *Service) Get() ([]estoquechave.Chave, error) {
	chave, err := s.Repository.Get()
	if err != nil {
		return nil, err
	}

	return chave, nil
}

func (s *Service) ServiceUpadete(chave *gormrepo.Chave) error {
	slog.Info("Preparando dados para salvar no banco de dados")
	quantatidade, err := estoquechave.NewChave(chave.Marca, chave.Modelo, chave.Quantidade)
	if err != nil {
		return fmt.Errorf("chave não informada")
	}

	slog.Info("Salvando no banco de dados")
	if err := s.Repository.Upadete(quantatidade); err != nil {
		return fmt.Errorf("atualizar chave: %w", err)
	}

	return nil
}

func (s *Service) ServiceGetId(id string) (*estoquechave.Chave, error) {
	if id == "" {
		return nil, fmt.Errorf("ID da chave não informado")
	}

	chave, err := s.Repository.GetKeyID(id)
	if err != nil {
		return nil, fmt.Errorf("consultar chave: %w", err)
	}

	return chave, nil
}

func (s *Service) ServiceDeleteID(id string) error {
	if id == "" {
		return fmt.Errorf("ID da chave não informado")
	}

	if err := s.Repository.DeleteID(id); err != nil {
		return fmt.Errorf("excluir chave: %w", err)
	}

	return nil
}
