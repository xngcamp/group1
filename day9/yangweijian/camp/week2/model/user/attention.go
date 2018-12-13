package user

import (
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) addAttention(uid, subUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	return c.UpdateId(
		uid,
		bson.M{"$push": bson.M{"sub_users": subUid}},
	)
}

func (userModel *UserModel) removeAttention(uid, unSubUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	return c.UpdateId(
		uid,
		bson.M{"$pull": bson.M{"sub_users": unSubUid}},
	)
}