package domain1

import "context"

// Repository is the domain1 storage
type Repository interface {
	// Registers inserts given Entity1 into storage
	Register(ctx context.Context, entity *Entity1) (entityID string, err error)

	// SetField3 sets field3 for Entity1
	SetField3(ctx context.Context, entityID string, field3 int) error
}
