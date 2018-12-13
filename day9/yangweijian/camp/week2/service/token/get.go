package token

import (
	"camp/week2/api"
	"camp/week2/model"
)

func (toeknService *TokenService) Get(tokenStr string) (tokenObj *api.Token, err error) {
	tokenModel := model.NewToken()

	tokenObj, err = tokenModel.Get(tokenStr)

	return
}