package token

import (
	"camp/week2/api"
	"camp/week2/mongo"
	"github.com/globalsign/mgo"
)

type TokenModel api.Token

func (t *TokenModel) Db() string {
	return "xngao"
}

func (t *TokenModel) Table() string {
	return "token"
}

func (t *TokenModel) GetC() (c *mgo.Collection) {
	db, table := t.Db(), t.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}