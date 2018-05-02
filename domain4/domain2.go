package domain4

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain2"
)

// NewDomain2Repository creates new domain2 repository implements by domain4
func NewDomain2Repository() domain2.Repository {
	return &domain2Repository{}
}

type domain2Repository struct{}

func (domain2Repository) Register(ctx context.Context, entity *domain2.Entity1) (string, error) {
	return "", nil
}
