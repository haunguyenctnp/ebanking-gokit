package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Uuid        string        `json:"uuid"  bson:"uuid"`
	Amount      float32       `json:"amount" bson:"amount"`
	Currency    string        `json:"currency" bson:"currency"`
	Date        string        `json:"date" bson:"date"`
	Description string        `json:"description" bson:"description"`
	Category    string        `json:"category" bson:"category"`
}

func (t Transaction) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
