package feed

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Update() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	fmt.Printf("%v\n", feed)
	err = c.Update(bson.M{"_id": feed.Id}, bson.M{"txt": feed.Txt})
	return
}