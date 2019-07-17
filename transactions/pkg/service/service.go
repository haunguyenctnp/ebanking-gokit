package service

import (
	"context"
	"github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/db"
	"github.com/haunguyenctnp/ebanking-gokit/transactions/pkg/io"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// TransactionsService describes the service.
type TransactionsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Get(ctx context.Context, uuid string) (t []io.Transaction, error error)
	Create(ctx context.Context, transaction io.Transaction) (t io.Transaction, error error)
	DeleteAll(ctx context.Context) (error error)
}

type basicTransactionsService struct{}

func (b *basicTransactionsService) Get(ctx context.Context, uuid string) (t []io.Transaction, error error) {
	// TODO implement the business logic of Get
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("transactions").C("transaction")
	error = c.Find(bson.M{"uuid": uuid}).All(&t)
	if t == nil {
		r := []io.Transaction{}
		return r, error
	}
	return t, error
}
func (b *basicTransactionsService) Create(ctx context.Context, transaction io.Transaction) (t io.Transaction, error error) {
	transaction.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	fmt.Printf("%+v\n", transaction)
	defer session.Close()
	c := session.DB("transactions").C("transaction")
	error = c.Insert(&transaction)
	return transaction, error
}

func (b *basicTransactionsService) DeleteAll(ctx context.Context) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("transactions").C("transaction")
	error = c.DropCollection()
	return error
}

// NewBasicTransactionsService returns a naive, stateless implementation of TransactionsService.
func NewBasicTransactionsService() TransactionsService {
	return &basicTransactionsService{}
}

// New returns a TransactionsService with all of the expected middleware wired in.
func New(middleware []Middleware) TransactionsService {
	svc := NewBasicTransactionsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
