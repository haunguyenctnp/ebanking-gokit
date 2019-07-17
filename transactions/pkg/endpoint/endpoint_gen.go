// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEndpoint       endpoint.Endpoint
	CreateEndpoint    endpoint.Endpoint
	DeleteAllEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.TransactionsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint:    MakeCreateEndpoint(s),
		DeleteAllEndpoint: MakeDeleteAllEndpoint(s),
		GetEndpoint:       MakeGetEndpoint(s),
	}
	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}
	for _, m := range mdw["DeleteAll"] {
		eps.DeleteAllEndpoint = m(eps.DeleteAllEndpoint)
	}
	return eps
}
