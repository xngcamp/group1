package feed

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type AddReq struct {
	Txt string `json:"txt"`
}

func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if addReq.Txt == "" {
		return
	}

	return true
}

type AddResp struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

// @prefilter("Cors")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi := api.NewFeed()
	feedApi.Txt = addReq.Txt
	feedApi.Id = bson.NewObjectId()
	if err := service.NewFeed().Add(feedApi); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := AddResp{feedApi.Id, feedApi.Txt}
	feed.ReplyOk(w, &resp)
}