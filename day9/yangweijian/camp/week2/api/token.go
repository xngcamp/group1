package api

import (
	"github.com/globalsign/mgo/bson"
)

type Token struct {
	UserId 	bson.ObjectId 	`bson:"uid"`
	Token 	string 			`bson:"token"`
	TimeOut int64			`bson:"time_out"`
}

func NewToken() *Token {
	return new(Token)
}