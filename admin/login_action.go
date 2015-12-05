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
		w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "验证码错误"}))
	} else {
		user := &data.User{Logger: logger}
		if user.Check(username, password) {
			s.Set("useradmin", username)
			w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success"}))
		} else {
			w.Write(jsonResponse(map[string]interface{}{"status": false, "msg": "用户名或密码错误"}))
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
	s.Delete("useradmin")
	w.Write(jsonResponse(map[string]interface{}{"status": true, "msg": "success"}))
}

func init() {
	ActionHandlers["login"] = login
	ActionHandlers["logout"] = logout
}
