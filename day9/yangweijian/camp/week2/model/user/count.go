package user

import (
	"github.com/globalsign/mgo/bson"
)

func (u *UserModel) Count(email string) (count int, err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	count, err = c.Find(bson.M{"email": email}).Count()
	return
}