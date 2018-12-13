package token

import (
	"camp/lib"
	"camp/week2/api"
	"time"
)

func (t *TokenModel) Add(user *api.User) (token string, err error) {
	c := t.GetC()
	defer c.Database.Session.Close()

	now := time.Now()
	str := now.String() + "-+-" + user.Password
	token = lib.HashSha256(str)
	aHour, _ := time.ParseDuration("24h")
	timeOut := now.Add(aHour)

	tokenObj := api.Token{user.Id, token, timeOut.Unix()}
	return token, c.Insert(tokenObj)
}