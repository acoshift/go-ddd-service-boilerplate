package mock

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
)

// Domain1Service mocks domain1 service
type Domain1Service struct {
	CreateFunc func(ctx context.Context, entity *domain1.Entity1) (entityID string, err error)
	UpdateFunc func(ctx context.Context, entity *domain1.Entity1) error
	Logic1Func func(ctx context.Context, arg1 string, arg2 int) error
	Logic2Func func(ctx context.Context, entityID string) error
}

// Create calls CreateFunc
func (s *Domain1Service) Create(ctx context.Context, entity *domain1.Entity1) (entityID string, err error) {
	return s.CreateFunc(ctx, entity)
}

// Update calls UpdateFunc
func (s *Domain1Service) Update(ctx context.Context, entity *domain1.Entity1) error {
	return s.UpdateFunc(ctx, entity)
}

// Logic1 calls Logic1Func
func (s *Domain1Service) Logic1(ctx context.Context, arg1 string, arg2 int) error {
	return s.Logic1Func(ctx, arg1, arg2)
}

// Logic2 calls Logic2Func
func (s *Domain1Service) Logic2(ctx context.Context, entityID string) error {
	return s.Logic2Func(ctx, entityID)
}
