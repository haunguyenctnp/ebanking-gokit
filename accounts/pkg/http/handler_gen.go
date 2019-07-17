// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/haunguyenctnp/ebanking-gokit/accounts/pkg/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeDepositHandler(m, endpoints, options["Deposit"])
	makeWithdrawHandler(m, endpoints, options["Withdraw"])
	makeDropHandler(m, endpoints, options["Drop"])
	makeCreateHandler(m, endpoints, options["Create"])
	makeGetAccountHandler(m, endpoints, options["GetAccount"])
	makeDepositAccountHandler(m, endpoints, options["DepositAccount"])
	makeWithdrawAccountHandler(m, endpoints, options["WithdrawAccount"])
	return m
}