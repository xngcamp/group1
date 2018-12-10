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

type UpdateReq struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

type UpdateResp struct {
}

// @prefilter("Cors")
func (feed *Feed) Update(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(feed.ReadBody(r), &updateReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi := api.NewFeed()
	feedApi.Id = updateReq.Id
	feedApi.Txt = updateReq.Txt
	if err := service.NewFeed().Update(feedApi); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := UpdateResp{}
	feed.ReplyOk(w, &resp)

	return
}