package model

import (
	"camp/week2/model/token"
	"camp/week2/model/user"
)

func NewUser() *user.UserModel {
	return new(user.UserModel)
}

func NewToken() *token.TokenModel {
	return new(token.TokenModel)
}