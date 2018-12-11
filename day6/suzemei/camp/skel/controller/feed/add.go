package feed

import (
	"camp/lib"
	"camp/skel/api"
	"camp/skel/service"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)



// AddReq 定义输入
type AddReq struct {
	//ID int64 `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if addReq.Txt=="" {
		return
	}

	ok = true
	return
}

// AddResp 定义输出
type AddResp struct {
	Txt string `json:"txt"`
}

// Add just for demo
// @postfilter("Boss")
func (feedc *FeedC) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"//打印日志

	var addReq *AddReq
	if err := json.Unmarshal(feedc.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feedc.ReplyFail(w, lib.CodePara)
		return
	}

	feedcApi := api.NewFeed()
	feedcApi.Txt = addReq.Txt
	if err := service.NewFeed().Add(feedcApi); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		feedc.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{}
	feedc.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedcApi, lib.ADD, nil)

	return
}
