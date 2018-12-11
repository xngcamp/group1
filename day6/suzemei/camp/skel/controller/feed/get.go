package feed

import (
	"encoding/json"
	"net/http"

	"camp/lib"
	"camp/skel/api"
	"camp/skel/service"

	"github.com/simplejia/clog/api"
)

// GetReq 定义输入
type GetReq struct {
	ID int `json:"id"`
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (getReq *GetReq) Regular() (ok bool) {
	if getReq == nil {
		return
	}

	if getReq.ID <= 0 ||getReq.Txt==""{
		return
	}

	ok = true
	return
}

// GetResp 定义输出
type GetResp struct {
	Feed *api.Feed `json:"skel,omitempty"`
}

// Get just for demo
// @postfilter("Boss")
func (feedc *FeedC) Get(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Get"

	var getReq *GetReq
	if err := json.Unmarshal(feedc.ReadBody(r), &getReq); err != nil || !getReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, getReq)
		feedc.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi, err := service.NewFeed().Get(getReq.ID,getReq.Txt)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, getReq)
		feedc.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetResp{
		Feed: feedApi,
	}
	feedc.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.GET, nil)

	return
}

