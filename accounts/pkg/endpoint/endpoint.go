package endpoint

import (
	io "github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/io"
	service "github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	io.Account
}

type CreateResponse struct {
	T     io.Account `json:"t"`
	Error error      `json:"error"`
}

func MakeCreateEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		t, error := s.Create(ctx, req.Account)
		return CreateResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

func (r CreateResponse) Failed() error {
	return r.Error
}

type DepositRequest struct{}

type DepositResponse struct {
	Error error `json:"error"`
}

func MakeDepositEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		error := s.Deposit(ctx)
		return DepositResponse{Error: error}, nil
	}
}

func (r DepositResponse) Failed() error {
	return r.Error
}

type WithdrawRequest struct{}

type WithdrawResponse struct {
	Error error `json:"error"`
}

func MakeWithdrawEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		error := s.Withdraw(ctx)
		return WithdrawResponse{Error: error}, nil
	}
}

func (r WithdrawResponse) Failed() error {
	return r.Error
}

type DropRequest struct{}

type DropResponse struct {
	Error error `json:"error"`
}

func MakeDropEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		error := s.Drop(ctx)
		return DropResponse{Error: error}, nil
	}
}

func (r DropResponse) Failed() error {
	return r.Error
}

type Failure interface {
	Failed() error
}

func (e Endpoints) Create(ctx context.Context, account io.Account) (t io.Account, error error) {
	request := CreateRequest{Account: account}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).T, response.(CreateResponse).Error
}

func (e Endpoints) Deposit(ctx context.Context) (error error) {
	request := DepositRequest{}
	response, err := e.DepositEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DepositResponse).Error
}

func (e Endpoints) Withdraw(ctx context.Context) (error error) {
	request := WithdrawRequest{}
	response, err := e.WithdrawEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WithdrawResponse).Error
}

func (e Endpoints) Drop(ctx context.Context) (error error) {
	request := DropRequest{}
	response, err := e.DropEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DropResponse).Error
}

type GetAccountRequest struct {
	Uuid string `json:"uuid"`
}

type GetAccountResponse []io.Account

func MakeGetAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAccountRequest)
		t, nil := s.GetAccount(ctx, req.Uuid)
		return t, nil
	}
}

func (r GetAccountResponse) Failed() error {
	return nil
}

func (e Endpoints) GetAccount(ctx context.Context, uuid string) (t []io.Account, error error) {
	request := GetAccountRequest{Uuid: uuid}
	response, err := e.GetAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAccountResponse), nil
}

type DepositAccountRequest struct {
	io.Deposit
}

type DepositAccountResponse io.Account

func MakeDepositAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DepositAccountRequest)
		t, nil := s.DepositAccount(ctx, req.Deposit)
		return t, nil
	}
}

func (r DepositAccountResponse) Failed() error {
	return nil
}

func (e Endpoints) DepositAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	request := DepositAccountRequest{Deposit: deposit}
	response, err := e.DepositAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Account), nil
}

type WithdrawAccountRequest struct {
	io.Deposit
}

type WithdrawAccountResponse struct {
	io.Account
}

func MakeWithdrawAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WithdrawAccountRequest)
		t, nil := s.WithdrawAccount(ctx, req.Deposit)
		return t, nil
	}
}

func (r WithdrawAccountResponse) Failed() error {
	return nil
}

func (e Endpoints) WithdrawAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	request := WithdrawAccountRequest{Deposit: deposit}
	response, err := e.WithdrawAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Account), nil
}
