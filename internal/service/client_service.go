package service

import (
	"context"
	"fmt"
	"storeapi/internal/data/repository"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type clientService struct {
	repo repository.ClientRepo
}

func NewClientService(repo repository.ClientRepo) ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) GetAll(ctx context.Context) ([]*models.Client, error) {
	list, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}
	return list, nil
}

func (s *clientService) GetById(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	model, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	return model, nil
}

func (s *clientService) GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error) {
	client, err := s.repo.GetByName(ctx, fullname)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	return client, nil
}

func (s *clientService) Create(ctx context.Context, model *models.Client) error {
	err := s.repo.Create(ctx, model)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	return nil
}

func (s *clientService) Update(ctx context.Context, model *models.Client) error {
	err := s.repo.Update(ctx, model)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			return err
		}
		return fmt.Errorf("failed to update client: %w", err)
	}
	return nil
}

func (s *clientService) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if err.Error() == "NO_AFFECTED" {
			return err
		}
		return fmt.Errorf("failed to delete client: %w", err)
	}
	return nil
}
