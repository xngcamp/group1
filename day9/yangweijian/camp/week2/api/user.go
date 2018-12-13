package api

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id 			bson.ObjectId 		`json:"id" bson:"_id"`
	Nick 		string 				`json:"nick" bson:"nick"`
	Sex 		byte				`json:"sex" bson:"sex"`
	Email 		string 				`json:"email" bson:"email"`
	Password 	string 				`json:"password" bson:"password"`
	FansUsers	[]*bson.ObjectId	`json:"fans_users" bson:"fans_users"`
	SubUsers	[]*bson.ObjectId	`json:"sub_users" bson:"sub_users"`
}

func NewUser() *User {
	return new(User)
}