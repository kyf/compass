package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	"html/template"
	"log"
	"net/http"
)

const (
	SERVER_ADDR string = ":9969"
)

func showTemplate(pathes []string, m *martini.ClassicMartini) {
	for _, p := range pathes {
		m.Get(fmt.Sprintf("/%s", p), (func(p string) func(http.ResponseWriter, *log.Logger) {
			return func(w http.ResponseWriter, logger *log.Logger) {
				var t *template.Template
				t, err := template.ParseFiles(fmt.Sprintf("../../admin/static/html/%s.html", p))
				if err != nil {
					log.Fatal(err)
				}
				w.Header().Set("content-type", "text/html")
				t.Execute(w, nil)
			}
		})(p))
	}
}

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()

	m.Get("/", func(w http.ResponseWriter) string {
		w.Header().Set("content-type", "text/html")
		return "<script type='text/javascript'>window.location.href='/login'</script>"
	})

	m.Use(martini.Static("../../admin/static/"))

	m.Get("/checkcode", func(r *http.Request, w http.ResponseWriter) {
		captcha.WriteImage(w, captcha.New(), 110, 40)
	})

	showTemplate([]string{"footer", "form", "index", "left", "login", "main", "right", "top"}, m)

	m.RunOnAddr(SERVER_ADDR)
}
