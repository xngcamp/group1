package feed

import (
	"camp/feed/api"
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Get() (feedRet *Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id": feed.Id}).One(&feedRet)
	return
}

func (feed *Feed) GetAll() (feedRet []*api.Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{}).All(&feedRet)
	return
}