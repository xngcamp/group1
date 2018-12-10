package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) Update(newFeed *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedModel.Id = newFeed.Id
	feedModel.Txt = newFeed.Txt

	err = feedModel.Update()
	return
}