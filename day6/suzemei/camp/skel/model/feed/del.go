package feed

import "github.com/globalsign/mgo"

// Del 定义删除操作
func (feed *Feed) Del() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Remove(feed)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}

