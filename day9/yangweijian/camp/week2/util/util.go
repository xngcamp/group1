package util

import (
	"camp/lib"
	"errors"
	"github.com/globalsign/mgo/bson"
)

func TokenToUid(base lib.IBase) (uid bson.ObjectId, err error) {
	if v, exists := base.GetParam("uid"); !exists {
		return bson.ObjectIdHex(""), errors.New("uid not exists")
	} else {
		uid = v.(bson.ObjectId)
	}
	return
}