package feed

import (
	"camp/feed/model"
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Del(id bson.ObjectId) (err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id

	err = feedModel.Del()

	return
}