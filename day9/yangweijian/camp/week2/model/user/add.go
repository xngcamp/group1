package user

import "camp/week2/api"

func (u *UserModel) Add(userApi *api.User) (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()

	return c.Insert(userApi)
}