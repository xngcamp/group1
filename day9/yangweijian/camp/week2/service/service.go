package service

import (
	"camp/week2/service/token"
	"camp/week2/service/user"
)

func NewUserService() *user.UserService {
	return new(user.UserService)
}

func NewTokenService() *token.TokenService {
	return new(token.TokenService)
}