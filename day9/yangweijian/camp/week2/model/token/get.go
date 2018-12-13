package token

import (
	"camp/week2/api"
	"github.com/globalsign/mgo/bson"
)

func (t *TokenModel) Get(token string) (tokenObj *api.Token, err error) {
	tokenObj = &api.Token{}
	c := t.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"token": token}).One(&tokenObj)
	return
}