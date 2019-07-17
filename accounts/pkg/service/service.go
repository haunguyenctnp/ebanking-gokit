package service

import (
	"github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/db"
	"github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/io"
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AccountsService describes the service.
type AccountsService interface {
	Deposit(ctx context.Context) (error error)
	Withdraw(ctx context.Context) (error error)
	Drop(ctx context.Context) (error error)

	Create(ctx context.Context, account io.Account) (t io.Account, error error)
	GetAccount(ctx context.Context, uuid string) (t []io.Account, error error)
	DepositAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error)
	WithdrawAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error)
}

type basicAccountsService struct{}

func (b *basicAccountsService) Create(ctx context.Context, account io.Account) (t io.Account, error error) {
	account.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	fmt.Printf("%+v\n", account)
	//
	var number int = int(rand.Float32() * 900000)
	var balance float32 = 0
	//
	if account.Type == "current" {
		balance = 5000
	}
	if account.Type == "credit" {
		balance = 40000
	}
	account.Number = number
	account.Balance = balance
	defer session.Close()
	c := session.DB("accounts").C("account")
	error = c.Insert(&account)
	return account, error
}

func (b *basicAccountsService) Deposit(ctx context.Context) (error error) {
	// TODO implement the business logic of Deposit
	return error
}
func (b *basicAccountsService) Withdraw(ctx context.Context) (error error) {
	// TODO implement the business logic of Withdraw
	return error
}
func (b *basicAccountsService) Drop(ctx context.Context) (error error) {
	// TODO implement the business logic of Drop
	return error
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService() AccountsService {
	return &basicAccountsService{}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountsService {
	var svc AccountsService = NewBasicAccountsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAccountsService) GetAccount(ctx context.Context, uuid string) (t []io.Account, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("accounts").C("account")
	error = c.Find(bson.M{"uuid": uuid}).All(&t)
	if t == nil {
		r := []io.Account{}
		return r, error
	}
	return t, error
}

func (b *basicAccountsService) DepositAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	//
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("accounts").C("account")
	value, err := strconv.ParseFloat(deposit.Amount, 32)

	var amount float32 = float32(value)
	// var number int = int(i1)
	// var numbers = make([]int, len(deposit.Number))
	i1, err := strconv.Atoi(deposit.Number[0])
	var number int = int(i1)
	// for index, element := range deposit.Number {
	// 	i1, err := strconv.Atoi(element)
	// 	_ = err

	// 	numbers[index] = number
	// }
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"balance": amount}},
		ReturnNew: true,
	}
	c.Find(bson.M{"number": number}).Apply(change, &t)
	return t, error
}

func (b *basicAccountsService) WithdrawAccount(ctx context.Context, deposit io.Deposit) (t io.Account, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("accounts").C("account")
	value, err := strconv.ParseFloat(deposit.Amount, 32)
	var amount float32 = float32(value)
	i1, err := strconv.Atoi(deposit.Number[0])
	var number int = int(i1)
	if err != nil {
		// do something sensible
	}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"balance": -amount}},
		ReturnNew: true,
	}
	c.Find(bson.M{"number": number}).Apply(change, &t)
	return t, error
}
