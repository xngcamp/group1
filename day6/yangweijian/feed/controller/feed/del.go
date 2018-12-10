package feed

import (
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type DelReq struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}

	return true
}

type DelResp struct {
}

func (feed *Feed) Del(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi, err := service.NewFeed().Get(delReq.Id)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	if feedApi == nil {
		detail := "feed not found"
		feed.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	if err := service.NewFeed().Del(delReq.Id); err != nil {
		clog.Error("%s feed.Del err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	feed.ReplyOk(w, resp)

	return
}