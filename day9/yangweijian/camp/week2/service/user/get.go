package user

import (
	"camp/week2/api"
	"camp/week2/model"
)

func (u *UserService) Get(user *api.User) (userRes *api.User, err error) {
	userModel := model.NewUser()
	userRes, err = userModel.Get(user)

	return
}