package user

import (
	"camp/lib"
	"camp/week2/api"
	"camp/week2/service"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

// 前端传过来的 user_info 参数
type UpdateInfo struct {
	Id 			bson.ObjectId 	`json:"id"`
	Nick 		string 			`json:"nick"`
	Sex 		byte 			`json:"sex"`
	Email 		string 			`json:"email"`
	Password 	string 			`json:"password"`
}

type UpdateReq struct {
	Info *UpdateInfo		`json:"update_info"`
}

func (u *UpdateReq) Regular() (ok bool) {
	if u == nil || u.Info == nil {
		return true
	}

	return
}

// @prefilter("All")
// @postfilter("Boss")
func (userController *User) Update(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(userController.ReadBody(r), &updateReq); err != nil || updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUser()
	userApi.Id 			= updateReq.Info.Id
	userApi.Nick 		= updateReq.Info.Nick
	userApi.Sex 		= updateReq.Info.Sex
	userApi.Email 		= updateReq.Info.Email
	userApi.Password 	= updateReq.Info.Password
	
	userService := service.NewUserService()
	if err := userService.Update(userApi); err != nil {
		clog.Error("%s update err: %v, req: %v", fun, err, updateReq)
		userController.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	//======================================
	// 暂时没有说返回值是什么，先不写
	var none interface{}
	userController.ReplyOk(w, none)
	//======================================
}