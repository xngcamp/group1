package feed

import (
	"camp/skel/api"
	"camp/skel/controller/feed"
	"camp/skel/model"
)

// Get 定义获取操作
func (feeds *Feeds) Get(id int,txt string) (feedApi *api.Feed, err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id
	feedModel.Txt = txt
	if feedModel, err = feedModel.Get(); err != nil {
		return   
	}

	if feedModel == nil {
		return
	}

	feedApi = (*api.Feed{})(feedModel)

	return
}

