package user

import (
	"camp/week2/api"
	"camp/week2/mongo"
	"github.com/globalsign/mgo"
)

type UserModel api.User

func (u *UserModel) Db() string {
	return "xngao"
}

func (u *UserModel) Table() string {
	return "user"
}

func (u *UserModel) GetC() (c *mgo.Collection) {
	db, table := u.Db(), u.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}