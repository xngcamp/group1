package feed

import (
	"camp/feed/api"
	"camp/feed/mongo"
	"github.com/globalsign/mgo"
)

type Feed api.Feed

func (feed *Feed) Db() (db string) {
	return "skel"
}

func (feed *Feed) Table() (table string)  {
	return "feed"
}

func (feed *Feed) GetC() (c *mgo.Collection) {
	db, table := feed.Db(), feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}