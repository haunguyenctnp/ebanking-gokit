package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Uuid     string        `json:"uuid"  bson:"uuid"`
	Email    string        `json:"email" bson:"email"`
	Phone    string        `json:"phone" bson:"phone"`
	Gender   string        `json:"gender" bson:"gender"`
	Dob      string        `json:"dob" bson:"dob"`
	Eid      string        `json:"eid" bson:"eid"`
	Password string        `json:"password" bson:"password"`
}

func (t User) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
