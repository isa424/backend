package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Email    string        `json:"email" bson:"email"`
		Username string        `json:"username"`
		Password string        `json:"password,omitempty" bson:"password"`
	}
)
