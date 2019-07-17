package service

import (
	"context"
	io "github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TransactionsService) TransactionsService

type loggingMiddleware struct {
	logger log.Logger
	next   TransactionsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TransactionsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TransactionsService) TransactionsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context, uuid string) (t []io.Transaction, error error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "error", error)
	}()
	return l.next.Get(ctx, uuid)
}
func (l loggingMiddleware) Create(ctx context.Context, transaction io.Transaction) (t io.Transaction, error error) {
	defer func() {
		l.logger.Log("method", "Create", "transaction", transaction, "t", t, "error", error)
	}()
	return l.next.Create(ctx, transaction)
}

func (l loggingMiddleware) DeleteAll(ctx context.Context) (error error) {
	defer func() {
		l.logger.Log("method", "DeleteAll", "error", error)
	}()
	return l.next.DeleteAll(ctx)
}
