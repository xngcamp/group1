package filter

import (
	"camp/week2/controller/user"
	"camp/week2/service"
	"github.com/simplejia/clog/api"
	"net/http"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	fun := "filter auth"

	// 获取 cookie 中的 token
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		clog.Error("%s err cookie: %v", fun, err)
	}
	tokenStr := tokenCookie.Value

	tokenService := service.NewTokenService()

	token, err := tokenService.Get(tokenStr)
	if err != nil {
		clog.Error("%s param err: %v", fun, err)
		return false
	}

	now := time.Now().Unix()
	if now < token.TimeOut {
		// 获取当前 token 的 uid
		if controller, exist := m["__C__"]; !exist {
			clog.Error("%s controller not exists: %v", fun, err)
			return false
		} else {
			c := controller.(*user.User)
			c.SetParam("uid", token.UserId)
			return true
		}
	} else {
		return false
	}
}