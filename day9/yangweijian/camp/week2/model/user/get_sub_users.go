package user

import (
	"camp/week2/api"
	"github.com/globalsign/mgo/bson"
)

func (userModel *UserModel) GetUsersByIds(uids []*bson.ObjectId) (users []*api.User, err error) {
	users = make([]*api.User, 0)

	c := userModel.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id": bson.M{"$in": uids}}).All(&users)
	return
}