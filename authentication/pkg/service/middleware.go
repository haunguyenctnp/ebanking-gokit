package service

import (
	io "github.com/haunguyenctnp/ebanking-gokit/authentication/pkg/io"
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthenticationService) AuthenticationService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthenticationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthenticationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthenticationService) AuthenticationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (t []io.User, error error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Create(ctx context.Context, user io.User) (t io.User, error error) {
	defer func() {
		l.logger.Log("method", "Create", "user", user, "t", t, "error", error)
	}()
	return l.next.Create(ctx, user)
}
func (l loggingMiddleware) Authenticate(ctx context.Context) (error error) {
	defer func() {
		l.logger.Log("method", "Authenticate", "error", error)
	}()
	return l.next.Authenticate(ctx)
}
