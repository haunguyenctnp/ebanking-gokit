package service

import (
	"context"

	io "github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AccountsService) AccountsService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountsService) AccountsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, account io.Account) (t io.Account, error error) {
	defer func() {
		l.logger.Log("method", "Create", "account", account, "t", t, "error", error)
	}()
	return l.next.Create(ctx, account)
}
func (l loggingMiddleware) Deposit(ctx context.Context) (error error) {
	defer func() {
		l.logger.Log("method", "Deposit", "error", error)
	}()
	return l.next.Deposit(ctx)
}
func (l loggingMiddleware) Withdraw(ctx context.Context) (error error) {
	defer func() {
		l.logger.Log("method", "Withdraw", "error", error)
	}()
	return l.next.Withdraw(ctx)
}
func (l loggingMiddleware) Drop(ctx context.Context) (error error) {
	defer func() {
		l.logger.Log("method", "Drop", "error", error)
	}()
	return l.next.Drop(ctx)
}

func (l loggingMiddleware) GetAccount(ctx context.Context, uuid string) (t []io.Account, error error) {
	defer func() {
		l.logger.Log("method", "GetAccount", "uuid", uuid, "t", t, "error", error)
	}()
	return l.next.GetAccount(ctx, uuid)
}

func (l loggingMiddleware) DepositAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	defer func() {
		l.logger.Log("method", "DepositAccount", "deposit", deposit, "t", t, "error", error)
	}()
	return l.next.DepositAccount(ctx, deposit)
}

func (l loggingMiddleware) WithdrawAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	defer func() {
		l.logger.Log("method", "WithdrawAccount", "deposit", deposit, "t", t, "error", error)
	}()
	return l.next.WithdrawAccount(ctx, deposit)
}
