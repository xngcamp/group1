package user

import (
	"camp/week2/api"
	"camp/week2/model"
	"errors"
)

func (u *UserService) Add(userApi *api.User) error {
	count, err := u.Count(userApi.Email)

	if count == 0 && err == nil {
		userModel := model.NewUser()
		return userModel.Add(userApi)
	} else if count != 0 {
		return errors.New("user has exists")
	}
	return err
}