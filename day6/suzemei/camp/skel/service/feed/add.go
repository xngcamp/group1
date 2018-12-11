package feed

import (
	"camp/skel/api"
	"camp/skel/model"
)

// Add 定义新增操作
func (feeds *Feeds) Add(feedApi *api.Feed) (err error) {
	feedsModel := model.NewFeed()
	//feedsModel.Id= feedApi.ID
	feedsModel.Txt= feedApi.Txt
	if err = feedsModel.Add(); err != nil {
		return
	}

	return
}