package user

import (
	"camp/lib"
	"camp/week2/api"
	"camp/week2/model"
)

func (userService *UserService) Update(userApi *api.User) error {
	userModel := model.NewUser()

	if userApi.Password != "" {
		userApi.Password = lib.HashSha256(userApi.Password)
	}

	return userModel.Update(userApi)
}