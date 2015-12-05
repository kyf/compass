package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	"github.com/kyf/compass/admin"
	"github.com/martini-contrib/sessions"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	SERVER_ADDR   string = ":9969"
	SCRIPT_LOGIN  string = "<script type='text/javascript'>window.location.href='/login'</script>"
	APP_STORE_DIR string = "/work/compass/app/"
)

func isLogin(w http.ResponseWriter, r *http.Request, s sessions.Session) bool {
	if strings.EqualFold("/login", r.URL.Path) {
		return true
	}

	if s.Get("useradmin") == nil {
		return false
	} else {
		return true
	}
}

func showTemplate(pathes []string, m *martini.ClassicMartini) {
	for _, p := range pathes {
		m.Get(fmt.Sprintf("/%s", p), (func(p string) func(http.ResponseWriter, *http.Request, sessions.Session, *log.Logger) {
			return func(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger) {
				w.Header().Set("content-type", "text/html")
				if ok := isLogin(w, r, s); !ok {
					w.Write([]byte(SCRIPT_LOGIN))
					return
				}
				var t *template.Template
				t, err := template.ParseFiles(fmt.Sprintf("../../admin/static/html/%s.html", p))
				if err != nil {
					log.Fatal(err)
				}
				t.Execute(w, nil)
			}
		})(p))
	}
}

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()

	m.Use(martini.Static("../../admin/static/"))
	m.Use(martini.Static(APP_STORE_DIR))
	m.Use(sessions.Sessions("compass_session", sessions.NewCookieStore([]byte("compass_session_cookie"))))

	m.Get("/", func(w http.ResponseWriter) string {
		w.Header().Set("content-type", "text/html")
		return SCRIPT_LOGIN
	})

	m.Get("/checkcode", func(r *http.Request, w http.ResponseWriter, s sessions.Session) {
		code := captcha.NewLen(4)
		s.Set("checkcode", code)
		captcha.WriteImage(w, code, 110, 40)
	})

	showTemplate([]string{"footer", "form", "index", "left", "login", "main", "right", "top"}, m)

	for actionName, actionHandler := range admin.ActionHandlers {
		m.Post(fmt.Sprintf("/action/%s", actionName), actionHandler)
	}

	m.RunOnAddr(SERVER_ADDR)
}
