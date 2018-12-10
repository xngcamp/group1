package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) Add(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedModel.Txt = feedApi.Txt
	feedModel.Id = feedApi.Id
	err = feedModel.Add()
	return
}