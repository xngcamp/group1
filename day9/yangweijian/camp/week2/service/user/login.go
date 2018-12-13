package user

import (
	"camp/lib"
	"camp/week2/api"
	tokenService "camp/week2/service/token"
	"errors"
	"github.com/simplejia/clog/api"
)

func (userService *UserService) Login(userApi *api.User, loginPassword string) (token string, err error) {
	userSql, err := userService.Get(userApi)
	if err != nil {
		return "", errors.New("email not right")
	}

	password := lib.HashSha256(loginPassword)
	if password != userSql.Password {
		clog.Info("user password not right")
		return "", errors.New("user password not right")
	}

	ts := &tokenService.TokenService{}
	token, err = ts.Add(userSql)
	if err != nil {
		return "", errors.New("create token err")
	}

	return
}