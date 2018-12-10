package feed

import (
	"camp/feed/api"
	"camp/feed/model"
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Get(id bson.ObjectId) (feedApi *api.Feed, err error)  {
	feedModel := model.NewFeed()
	feedModel.Id = id
	if feedModel, err = feedModel.Get(); err != nil {
		return
	}

	if feedModel == nil {
		return
	}

	feedApi = (*api.Feed)(feedModel)
	return
}

func (feed *Feed) GetAll() (feedApi []*api.Feed, err error) {
	feedModel := model.NewFeed()

	feedApi, err = feedModel.GetAll()
	return
}