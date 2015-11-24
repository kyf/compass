package main

import (
	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	"html/template"
	"log"
	"net/http"
)

const (
	SERVER_ADDR string = ":9969"
)

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Get("/", func() string {
		return "server is running ..."
	})
	m.Use(martini.Static("./static/"))
	m.Get("/checkcode", func(r *http.Request, w http.ResponseWriter) {
		captcha.WriteImage(w, captcha.New(), 110, 40)
	})
	m.Get("/login", func(w http.ResponseWriter, logger *log.Logger) {
		var t *template.Template
		t, err := template.ParseFiles("./static/html/login.html")
		if err != nil {
			log.Fatal(err)
		}
		values := map[string]template.HTML{"name": template.HTML("<tt>")}
		w.Header().Set("content-type", "text/html")
		t.Execute(w, values)
	})

	m.RunOnAddr(SERVER_ADDR)
}
