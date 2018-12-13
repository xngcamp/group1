package user

import (
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) UnSubUser(uid bson.ObjectId, unSubUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	err = userModel.removeAttention(uid, unSubUid)
	if err != nil {
		return
	}

	return userModel.removeFans(unSubUid, uid)
}