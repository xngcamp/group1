package feed

import (
	"camp/feed/service"
	"camp/lib"
	"github.com/simplejia/clog/api"
	"net/http"
)

// @prefilter("Cors")
func (feed *Feed) GetAll(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.GetAll"

	feeds, err := service.NewFeed().GetAll()
	if err != nil {
		clog.Error("%s feed.Add err: %v", fun, err)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	feed.ReplyOk(w, feeds)
}