package mock

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
)

// Domain1Repository is the mock struct for domain1 repository
type Domain1Repository struct {
	RegisterFunc  func(ctx context.Context, entity *domain1.Entity1) (string, error)
	SetField3Func func(ctx context.Context, entityID string, field3 int) error
}

// Register calls RegisterFunc
func (r *Domain1Repository) Register(ctx context.Context, entity *domain1.Entity1) (string, error) {
	return r.RegisterFunc(ctx, entity)
}

// SetField3 calls SetField3 func
func (r *Domain1Repository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return r.SetField3Func(ctx, entityID, field3)
}
