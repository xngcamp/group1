//go:generate wsp -s -d

package main

import (
	"fmt"
	"net/http"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/lc"
	"github.com/simplejia/utils"
	"camp/lib"
	"camp/week2/conf"
)

func init() {
	lc.Init(1e5)

	clog.AddrFunc = func() (string, error) {
		return lib.NameWrap(conf.C.Addrs.Clog)
	}
	clog.Init(conf.C.Clog.Name, "", conf.C.Clog.Level, conf.C.Clog.Mode)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clog.Error("%s is not found", r.RequestURI)
		http.NotFound(w, r)
	})
}

func main() {
	fun := "main"
	clog.Info(fun)

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.C.App.Port)
	err := utils.ListenAndServe(addr, nil)
	if err != nil {
		clog.Error("%s err: %v, addr: %v", fun, err, addr)
	}
}
