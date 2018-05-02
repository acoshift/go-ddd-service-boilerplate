package domain2

import (
	"github.com/acoshift/go-ddd-service-boilerplate/domain3"
)

// Entity1 type
type Entity1 struct {
	ID     string
	Field1 string
	Field2 domain3.Entity1
}
