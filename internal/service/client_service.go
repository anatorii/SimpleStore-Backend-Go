package service

import (
	"context"
	"storeapi/internal/domain/models"
	"storeapi/internal/domain/repos"

	"github.com/google/uuid"
)

type clientService struct {
	repo repos.ClientRepoInt
}

func NewClientService(repo repos.ClientRepoInt) ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) GetAll(ctx context.Context) ([]*models.Client, error) {
	return nil, nil
}

func (s *clientService) GetById(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	client, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *clientService) GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error) {
	client, err := s.repo.GetByName(ctx, fullname)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *clientService) Create(ctx context.Context, model *models.Client) error {
	return nil
}

func (s *clientService) Update(ctx context.Context, model *models.Client) error {
	return nil
}

func (s *clientService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
