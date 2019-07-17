package endpoint

import (
	io "github.com/haunguyenctnp/ebanking-gokit/authentication/pkg/io"
	service "github.com/haunguyenctnp/ebanking-gokit/authentication/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T     []io.User `json:"t"`
	Error error     `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, error := s.Get(ctx)
		return GetResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	User io.User `json:"user"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	T     io.User `json:"t"`
	Error error   `json:"error"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		t, error := s.Create(ctx, req.User)
		return CreateResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Error
}

// AuthenticateRequest collects the request parameters for the Authenticate method.
type AuthenticateRequest struct{}

// AuthenticateResponse collects the response parameters for the Authenticate method.
type AuthenticateResponse struct {
	Error error `json:"error"`
}

// MakeAuthenticateEndpoint returns an endpoint that invokes Authenticate on the service.
func MakeAuthenticateEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		error := s.Authenticate(ctx)
		return AuthenticateResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r AuthenticateResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.User, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, user io.User) (t io.User, error error) {
	request := CreateRequest{User: user}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).T, response.(CreateResponse).Error
}

// Authenticate implements Service. Primarily useful in a client.
func (e Endpoints) Authenticate(ctx context.Context) (error error) {
	request := AuthenticateRequest{}
	response, err := e.AuthenticateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AuthenticateResponse).Error
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T     []io.User `json:"t"`
	Error error     `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, error := s.Get(ctx)
		return GetResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.User, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Error
}
