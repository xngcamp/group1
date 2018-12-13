package user

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) addFans(uid, subUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	return c.UpdateId(
		uid,
		bson.M{"$push": bson.M{"fans_users": subUid}},
	)
}

func (userModel *UserModel) removeFans(uid, unSubUid bson.ObjectId) (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()
	fmt.Printf("fans: one: %s, two: %s\n", uid, unSubUid)

	return c.UpdateId(
		uid,
		bson.M{"$pull": bson.M{"fans_users": unSubUid}},
	)
}