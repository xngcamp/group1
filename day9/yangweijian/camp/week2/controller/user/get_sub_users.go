package user

import (
	"camp/lib"
	"camp/week2/service"
	"github.com/simplejia/clog/api"
	"net/http"
)

// @prefilter("All")
// @postfilter("Boss")
func (userController *User) GetSubUsers(w http.ResponseWriter, r *http.Request) {
	fun := "week2.userController.GetSubUsers"

	users, err := service.NewUserService().GetSubUsers(userController)
	if err != nil {
		clog.Error("%s param err: %v", fun, err)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}

	userController.ReplyOk(w, users)
}