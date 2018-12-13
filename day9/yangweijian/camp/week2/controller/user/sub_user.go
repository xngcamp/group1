package user

import (
	"camp/lib"
	"camp/week2/service"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type SubUserReq struct {
	Id bson.ObjectId 	`json:"need_sub_uid"`
}

func (s *SubUserReq) Regular() (ok bool) {
	if s == nil || s.Id.String() == "" {
		return true
	}

	return false
}

// @prefilter("All")
// @postfilter("Boss")
func (userController *User) SubUser(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.SubUser"

	var subUserReq *SubUserReq
	if err := json.Unmarshal(userController.ReadBody(r), &subUserReq); err != nil || subUserReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, subUserReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	if err := service.NewUserService().SubUser(userController, subUserReq.Id); err != nil {
		clog.Error("%s sub_user err: %v, req: %v", fun, err, subUserReq)
		userController.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	//======================================
	// 暂时没有说返回值是什么，先不写
	var none interface{}
	userController.ReplyOk(w, none)
	//======================================

	w.Header().Set("Access-Control-Allow-Origin", "*")
}