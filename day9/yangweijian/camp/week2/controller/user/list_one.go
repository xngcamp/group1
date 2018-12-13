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

type ListOneReq struct {
	Id bson.ObjectId 		`json:"uid"`
}

func (l *ListOneReq) Regular() (ok bool) {
	if l == nil || l.Id.String() == "" {
		return true
	}

	return false
}

type ListOneRes struct {
	Id 		bson.ObjectId 	`json:"id" bson:"_id"`
	Nick  	string	 		`json:"nick" bson:"nick"`
	Sex 	byte			`json:"sex" bson:"sex"`
	Email 	string 			`json:"email" bson:"email"`
}

// @prefilter("All")
// @postfilter("Boss")
func (userController *User) ListOne(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.ListOne"

	var listOneReq *ListOneReq
	if err := json.Unmarshal(userController.ReadBody(r), &listOneReq); err != nil || listOneReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, listOneReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUser()
	userApi.Id = listOneReq.Id

	userResp, err := service.NewUserService().Get(userApi)
	if err != nil {
		clog.Error("%s get user err: %v, req: %v", fun, err, listOneReq)
		userController.ReplyFailWithDetail(w, lib.CodeSrv, "email not right")
		return
	}

	resp := &ListOneRes{}
	resp.Id 		= userResp.Id
	resp.Email 		= userResp.Email
	resp.Nick 		= userResp.Nick
	resp.Sex 		= userResp.Sex

	userController.ReplyOk(w, resp)
}