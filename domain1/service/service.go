package service

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
)

// New creates new domain1 service
func New(repo domain1.Repository) domain1.Service {
	return &service{repo}
}

type service struct {
	repo domain1.Repository
}

func (s *service) Create(ctx context.Context, entity *domain1.Entity1) (entityID string, err error) {
	return s.repo.Register(ctx, entity)
}

func (s *service) Update(ctx context.Context, entity *domain1.Entity1) error {
	return nil
}

func (s *service) Logic1(ctx context.Context, arg1 string, arg2 int) error {
	return nil
}

func (s *service) Logic2(ctx context.Context, entityID string) error {
	return nil
}
