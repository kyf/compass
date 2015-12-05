package admin

import (
	"github.com/kyf/compass/data"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
	"strings"
)

func AdSetting(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	state := r.PostForm.Get("state")

	setting := &data.Setting{Logger: logger}
	if strings.EqualFold(state, "true") {
		setting.Write(0)
	} else {
		setting.Write(1)
	}

	w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success"}))
}

func GetSetting(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	setting := &data.Setting{Logger: logger}
	setting.Read()

	w.Write(jsonResponse(map[string]interface{}{"status": true, "data": setting}))
}

func init() {
	ActionHandlers["adsetting"] = AdSetting
	ActionHandlers["getsetting"] = GetSetting
}
