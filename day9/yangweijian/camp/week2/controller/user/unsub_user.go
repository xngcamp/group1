package user

import (
	"camp/lib"
	"camp/week2/service"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type UnSubUserReq struct {
	Id bson.ObjectId	`json:"need_unsub_uid"`
}

func (s *UnSubUserReq) Regular() (ok bool) {
	if s == nil || s.Id.String() == "" {
		return true
	}

	return false
}

// @prefilter("All")
// @postfilter("Boss")
func (userController *User) UnsubUser(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.UnsubUser"

	var unSubUserReq *UnSubUserReq
	if err := json.Unmarshal(userController.ReadBody(r), &unSubUserReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, unSubUserReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	if err := service.NewUserService().UnSubUser(userController, unSubUserReq.Id); err != nil {
		clog.Error("%s unsub_user err: %v, req: %v", fun, err, unSubUserReq)
		userController.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	//======================================
	// 暂时没有说返回值是什么，先不写
	var none interface{}
	userController.ReplyOk(w, none)
	//======================================
}