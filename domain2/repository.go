package domain2

import "context"

// Repository is the domain2 storage
type Repository interface {
	// Registers inserts given Entity1 into storage
	Register(ctx context.Context, entity *Entity1) (entityID string, err error)
}
