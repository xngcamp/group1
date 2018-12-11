package feed

import "camp/skel/model"

// Del 定义删除操作
func (feeds *Feeds) Del(id int,txt string) (err error) {
	feedsModel := model.NewFeed()
	feedsModel.Id= id
	feedsModel.Txt= txt

	if err = feedsModel.Del(); err != nil {
		return
	}

	return
}
