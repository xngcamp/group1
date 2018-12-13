package user

import (
	"camp/lib"
	"camp/week2/api"
	"camp/week2/model"
	"camp/week2/util"
)

func (userService *UserService) getUserByIBase(base lib.IBase) (user *api.User, err error) {
	// 获取 token 中的 uid
	uid, err := util.TokenToUid(base)
	if err != nil {
		return nil, err
	}

	userModel := model.NewUser()
	user, err = userModel.Get(&api.User{Id: uid})
	if err != nil {
		return nil, err
	}

	return
}


func (userService *UserService) GetSubUsers(base lib.IBase) (users []*api.User, err error) {
	user, err := userService.getUserByIBase(base)
	if err != nil {
		return nil, err
	}

	if len(user.SubUsers) == 0 {
		return make([]*api.User, 0), nil
	}

	userModel := model.NewUser()
	users, err = userModel.GetUsersByIds(user.SubUsers)
	for _, u := range users{
		u.Password = ""
	}

	return
}

func (userService *UserService) GetFansUsers(base lib.IBase) (users []*api.User, err error) {
	user, err := userService.getUserByIBase(base)
	if err != nil {
		return nil, err
	}

	if len(user.FansUsers) == 0 {
		return make([]*api.User, 0), nil
	}

	userModel := model.NewUser()
	users, err = userModel.GetUsersByIds(user.FansUsers)
	for _, u := range users{
		u.Password = ""
	}

	return
}