package user

import "camp/week2/model"

func (u *UserService) Count(email string) (count int, err error) {
	userModel := model.NewUser()
	count, err = userModel.Count(email)

	return
}