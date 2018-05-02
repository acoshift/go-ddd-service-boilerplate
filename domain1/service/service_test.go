package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
	"github.com/acoshift/go-ddd-service-boilerplate/domain1/service"
	"github.com/acoshift/go-ddd-service-boilerplate/mock"
)

func TestServiceCreate(t *testing.T) {
	ctx := mock.Context{}
	repo := mock.Domain1Repository{
		RegisterFunc: func(ctx context.Context, entity *domain1.Entity1) (string, error) {
			assert.NotNil(t, entity)
			return "abc", nil
		},
	}

	s := service.New(&repo)

	id, err := s.Create(&ctx, &domain1.Entity1{})
	assert.NoError(t, err)
	assert.Equal(t, "abc", id)
}
