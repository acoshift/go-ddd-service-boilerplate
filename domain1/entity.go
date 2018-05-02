package domain1

import (
	"github.com/acoshift/go-ddd-service-boilerplate/domain3"
)

// Entity1 type
type Entity1 struct {
	ID     string
	Field1 string
	Field2 Entity2
	Field3 int
	Field4 domain3.Entity1
}

// Entity2 type
type Entity2 struct {
	Field1 string
	Field2 bool
}
