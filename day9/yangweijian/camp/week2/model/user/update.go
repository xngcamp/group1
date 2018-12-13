package user

import (
	"camp/week2/api"
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) Update(userApi *api.User) error {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	q := bson.M{}
	if userApi.Password != "" {
		q["password"] = userApi.Password
	}
	if userApi.Sex == 1 || userApi.Sex == 2 {
		q["sex"] = userApi.Sex
	}
	if userApi.Email != "" {
		q["email"] = userApi.Email
	}
	if userApi.Nick != "" {
		q["nick"] = userApi.Nick
	}

	return c.UpdateId(userApi.Id, bson.M{"$set": q})
}