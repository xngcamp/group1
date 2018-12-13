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

type RegisterReq struct {
	Nick string 		`json:"nick"`
	Sex byte 			`json:"sex"`
	Email string 		`json:"email"`
	Password string 	`json:"password"`
}

func (u *RegisterReq) Regular() (ok bool) {
	if u == nil || u.Nick == "" || u.Email == "" ||u.Password == "" ||
		(u.Sex != 1 && u.Sex != 2){
		return true
	}

	return
}

type RegisterRes struct {
	Id bson.ObjectId 	`json:"id"`
	Nick string 		`json:"nick"`
	Sex byte 			`json:"sex"`
	Email string 		`json:"email"`
}

// @prefilter("Cors")
// @postfilter("Boss")
func (userController *User) Register (w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.Register"

	var registerReq *RegisterReq
	if err := json.Unmarshal(userController.ReadBody(r), &registerReq); err != nil || registerReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, registerReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUser()
	userApi.Id 			= bson.NewObjectId()
	userApi.Nick 		= registerReq.Nick
	userApi.Sex 		= registerReq.Sex
	userApi.Email 		= registerReq.Email
	userApi.Password 	= lib.HashSha256(registerReq.Password)

	userService := service.NewUserService()
	if err := userService.Add(userApi); err != nil {
		clog.Error("%s add err: %v, req: %v", fun, err, registerReq)
		userController.ReplyFailWithDetail(w, lib.CodeSrv, err.Error())
		return
	}

	resp := &RegisterRes{
		Id: 	userApi.Id,
		Nick: 	userApi.Nick,
		Sex: 	userApi.Sex,
		Email:	userApi.Email,
	}
	userController.ReplyOk(w, resp)
}