package domain4

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
)

// NewDomain1Repository creates domain1 repository implements by domain4
func NewDomain1Repository() domain1.Repository {
	return &domain1Repository{}
}

type domain1Repository struct{}

func (domain1Repository) Register(ctx context.Context, entity *domain1.Entity1) (string, error) {
	return "", nil
}

// SetField3 sets field3 for Entity1
func (domain1Repository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return nil
}
