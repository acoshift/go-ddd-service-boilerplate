package endpoint

import (
	"context"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
)

// New creates new domain1 endpoint
func New(s domain1.Service) domain1.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s domain1.Service
}

func (ep *endpoint) Create(ctx context.Context, req *domain1.CreateRequest) (*domain1.CreateResponse, error) {
	id, err := ep.s.Create(ctx, &domain1.Entity1{
		Field2: domain1.Entity2{
			Field1: req.Field1,
		},
	})
	if err != nil {
		return nil, err
	}
	return &domain1.CreateResponse{ID: id}, nil
}
