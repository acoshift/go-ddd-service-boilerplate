package endpoint_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
	"github.com/acoshift/go-ddd-service-boilerplate/domain1/endpoint"
	"github.com/acoshift/go-ddd-service-boilerplate/mock"
)

func TestCreate(t *testing.T) {
	ctx := mock.Context{}
	s := mock.Domain1Service{
		CreateFunc: func(ctx context.Context, entity *domain1.Entity1) (entityID string, err error) {
			assert.Equal(t, "field1_data", entity.Field2.Field1)
			return "abc", nil
		},
	}

	ep := endpoint.New(&s)
	resp, err := ep.Create(&ctx, &domain1.CreateRequest{
		Field1: "field1_data",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "abc", resp.ID)
}
