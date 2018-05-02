package domain4_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
	"github.com/acoshift/go-ddd-service-boilerplate/domain4"
	"github.com/acoshift/go-ddd-service-boilerplate/mock"
)

func TestDomain1RepositoryRegister(t *testing.T) {
	// integrate test with real database

	// init database
	// defer cleanup database

	repo := domain4.NewDomain1Repository()
	ctx := mock.Context{}
	// inject db to context
	id, err := repo.Register(&ctx, &domain1.Entity1{})
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}
