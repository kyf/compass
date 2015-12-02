package admin

import (
	"fmt"
	"github.com/dchest/captcha"
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
	if !captcha.VerifyString(checkcode, code.(string)) {
		w.Write([]byte("check is wrong"))
	} else {
		w.Write([]byte("check right"))
	}
	w.Write([]byte(fmt.Sprintf("%s, %s", username, password)))
}

func logout(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {

}

func init() {
	ActionHandlers["login"] = login
	ActionHandlers["logout"] = logout
}
