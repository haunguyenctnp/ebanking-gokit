package service

import (
	"github.com/haunguyenctnp/ebanking-gokit/authentication/pkg/io"
	"context"
)

// AuthenticationService describes the service.
type AuthenticationService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Get(ctx context.Context) (t []io.User, error error)
	Create(ctx context.Context, user io.User) (t io.User, error error)
	Authenticate(ctx context.Context) (error error)
}

type basicAuthenticationService struct{}

func (b *basicAuthenticationService) Get(ctx context.Context) (t []io.User, error error) {
	// TODO implement the business logic of Get
	return t, error
}
func (b *basicAuthenticationService) Create(ctx context.Context, user io.User) (t io.User, error error) {
	// TODO implement the business logic of Create
	return t, error
}
func (b *basicAuthenticationService) Authenticate(ctx context.Context) (error error) {
	// TODO implement the business logic of Authenticate
	return error
}

// NewBasicAuthenticationService returns a naive, stateless implementation of AuthenticationService.
func NewBasicAuthenticationService() AuthenticationService {
	return &basicAuthenticationService{}
}

// New returns a AuthenticationService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthenticationService {
	var svc AuthenticationService = NewBasicAuthenticationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
