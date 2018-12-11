package feed

// Add 定义新增操作
func (feed *Feed) Add() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(feed)
	if err != nil {
		return
	}

	return
}

