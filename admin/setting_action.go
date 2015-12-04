package admin

import (
	"github.com/kyf/compass/data"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

func AdSetting(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	state := r.PostForm.Get("state")

	setting := &data.Setting{logger: logger}
	logger.Printf("state is %v, setting is %v", setting)
	w.Write([]byte("ss"))
}

func init() {
	ActionHandlers["adsetting"] = AdSetting
}
