package user

import (
	"camp/week2/api"
	"github.com/globalsign/mgo/bson"
)

func (u *UserModel) Get(user *api.User) (userRes *api.User, err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	q := bson.M{}

	if user.Id != "" {
		q["_id"] = user.Id
	}
	if user.Email != "" {
		q["email"] = user.Email
	}

	err = c.Find(q).One(&userRes)
	return
}