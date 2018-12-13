package user

import (
	"camp/lib"
	"camp/week2/api"
	"camp/week2/service"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type LoginReq struct {
	Email string		`json:"email"`
	Password string		`json:"password"`
}

func (l *LoginReq) Register() (ok bool) {
	if l == nil || l.Email == "" || l.Password == "" {
		return true
	}

	return
}

type LoginRes struct {
	Token string 		`json:"token"`
}

// @prefilter("Cors")
// @postfilter("Boss")
func (userController *User) Login(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.Login"
	
	var loginReq *LoginReq
	if err := json.Unmarshal(userController.ReadBody(r), &loginReq); err != nil || loginReq.Register() {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUser()
	userApi.Email = loginReq.Email

	token, err := service.NewUserService().Login(userApi, loginReq.Password)
	if err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFailWithDetail(w, lib.CodePara, err.Error())
		return
	}

	resp := &LoginRes{token}
	userController.ReplyOk(w, resp)
}