package admin

import (
	"github.com/dchest/captcha"
	"github.com/kyf/compass/data"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	checkcode := r.PostForm.Get("checkcode")

	code := s.Get("checkcode")
	if !captcha.VerifyString(code.(string), checkcode) {
		w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "checkcode is wrong"}))
	} else {
		user := &data.User{}
		if user.Check(username, password) {
			w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success"}))
		} else {
			w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "username or password is wrong"}))
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {

}

func init() {
	ActionHandlers["login"] = login
	ActionHandlers["logout"] = logout
}
