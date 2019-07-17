package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Uuid     string        `json:"uuid"  bson:"uuid"`
	Type     string        `json:"type" bson:"type"`
	Currency string        `json:"currency" bson:"currency"`
	Balance  float32       `json:"balance" bson:"balance"`
	Number   int           `json:"number" bson:"number"`
}

type Deposit struct {
	Number []string `json:"number"  bson:"number"`
	Amount string   `json:"amount"  bson:"amount"`
}

func (t Account) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}

func (t Deposit) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
