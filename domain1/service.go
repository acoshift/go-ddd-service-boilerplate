package domain1

import "context"

// Service is the domain1 service
type Service interface {
	// Create creates new Entity1
	Create(ctx context.Context, entity *Entity1) (entityID string, err error)

	// Update updates Entity1
	Update(ctx context.Context, entity *Entity1) error

	// Logic1 does logic1
	Logic1(ctx context.Context, arg1 string, arg2 int) error

	// Logic3 does logic2
	Logic2(ctx context.Context, entityID string) error
}
