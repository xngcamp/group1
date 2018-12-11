package feed

import (
	"camp/skel/mongo"
	"github.com/globalsign/mgo"
)
type Feed struct {
	Id int `json:"id"`
	Txt string `json:"txt"`
}
//
// Db 返回db name
func (feed *Feed) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (feed *Feed) Table() (table string) {
	return "feed"
}

// GetC 返回db col
func (feed *Feed) GetC() (c *mgo.Collection) {
	db, table := feed.Db(), feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
