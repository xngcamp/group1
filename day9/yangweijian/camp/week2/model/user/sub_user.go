package user

import (
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) SubUser(uid, subUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	err = userModel.addAttention(uid, subUid)
	if err != nil {
		return
	}

	return userModel.addFans(subUid, uid)
}