package api

import "github.com/globalsign/mgo/bson"

type Feed struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

func NewFeed() *Feed {
	return new(Feed)
}