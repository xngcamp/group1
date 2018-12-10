package feed

import "github.com/globalsign/mgo/bson"

func (feed *Feed) Del() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Remove(bson.M{"_id": feed.Id})
	return
}