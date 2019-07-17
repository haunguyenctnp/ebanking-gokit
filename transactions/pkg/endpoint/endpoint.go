package endpoint

import (
	"context"
	io "github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/io"
	service "github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	UUID string `json:"uuid"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse []io.Transaction

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		t, error := s.Get(ctx, req.UUID)
		return t, error
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return nil
}

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	io.Transaction
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse io.Transaction

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		t, error := s.Create(ctx, req.Transaction)
		return t, error
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return nil
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, uuid string) (t []io.Transaction, error error) {
	request := GetRequest{UUID: uuid}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse), err
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, transaction io.Transaction) (t CreateResponse, error error) {
	request := CreateRequest{Transaction: transaction}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse), err
}

// DeleteAllRequest struct
type DeleteAllRequest struct{}

// DeleteAllResponse struct
type DeleteAllResponse struct {
	Error error `json:"error"`
}

// MakeDeleteAllEndpoint returns an endpoint that invokes DeleteAll on the service.
func MakeDeleteAllEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		error := s.DeleteAll(ctx)
		return DeleteAllResponse{Error: error}, nil
	}
}

// Failed reponse
func (r DeleteAllResponse) Failed() error {
	return r.Error
}

// DeleteAll implements Service
func (e Endpoints) DeleteAll(ctx context.Context) (error error) {
	request := DeleteAllRequest{}
	response, err := e.DeleteAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAllResponse).Error
}
