package token

import (
	"camp/week2/api"
	"camp/week2/model"
)

func (tokenService *TokenService) Add(user *api.User) (token string, err error) {
	tokenModel := model.NewToken()

	return tokenModel.Add(user)
}