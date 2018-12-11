package feed

import (
	"encoding/json"
	"net/http"

	"camp/lib"
	"camp/skel/service"

	"github.com/simplejia/clog/api"
)

// DelReq 定义输入
type DelReq struct {
	ID int `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}

	if delReq.ID <= 0 ||delReq.Txt==""{
		return
	}

	ok = true
	return
}

// DelResp 定义输出
type DelResp struct {
}

// Del just for demo
// @postfilter("Boss")
func (feedc *FeedC) Del(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feedc.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feedc.ReplyFail(w, lib.CodePara)
		return
	}

	feedcApi, err := service.NewFeed().Get(delReq.ID,delReq.Txt)
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, delReq)
		feedc.ReplyFail(w, lib.CodeSrv)
		return
	}

	if feedcApi == nil {
		detail := "skel not found"
		feedc.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	if err := service.NewFeed().Del(delReq.ID,delReq.Txt); err != nil {
		clog.Error("%s skel.Del err: %v, req: %v", fun, err, delReq)
		feedc.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	feedc.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedcApi, lib.DELETE, nil)

	return
}

