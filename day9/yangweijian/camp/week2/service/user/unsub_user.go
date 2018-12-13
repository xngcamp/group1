package user

import (
	"camp/lib"
	"camp/week2/model"
	"camp/week2/util"
	"github.com/globalsign/mgo/bson"
)

func (userService *UserService) UnSubUser(base lib.IBase, unSubUid bson.ObjectId) (err error) {
	// 获取 token 中的 uid
	uid, err := util.TokenToUid(base)
	if err != nil {
		return
	}

	return model.NewUser().UnSubUser(uid, unSubUid)
}