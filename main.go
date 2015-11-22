package main

import (
	//"bytes"
	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	//"image"
	//"image/jpeg"
	"net/http"
	//"time"
)

const (
	SERVER_ADDR string = ":9969"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "server is running ..."
	})
	m.Use(martini.Static("./static/"))
	m.Get("/checkcode", func(r *http.Request, w http.ResponseWriter) {
		//http.ServeContent(w, r, name, time.Now(), bytes.NewReader(buf.Bytes()))
		captcha.WriteImage(w, captcha.New(), 110, 40)
	})

	m.RunOnAddr(SERVER_ADDR)
}
