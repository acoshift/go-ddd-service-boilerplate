package main

import (
	"net/http"

	"github.com/acoshift/go-ddd-service-boilerplate/domain1"
	domain1endpoint "github.com/acoshift/go-ddd-service-boilerplate/domain1/endpoint"
	domain1handler "github.com/acoshift/go-ddd-service-boilerplate/domain1/handler"
	domain1service "github.com/acoshift/go-ddd-service-boilerplate/domain1/service"
	"github.com/acoshift/go-ddd-service-boilerplate/domain4"
)

func main() {
	// init repos
	domain1Repo := domain4.NewDomain1Repository()

	// init services
	domain1Service := domain1service.New(domain1Repo)

	// init endpoints
	domain1Endpoint := domain1endpoint.New(domain1Service)

	mux := http.NewServeMux()
	mux.Handle("/", domain1handler.New(domain1Service))
	mux.Handle("/domain1/", http.StripPrefix("/domain1", domain1.NewHTTPTransport(domain1Endpoint)))

	http.ListenAndServe(":8080", mux)
}
