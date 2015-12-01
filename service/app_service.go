package service

import (
	"github.com/kyf/compass/data"
	"log"
	"net/http"
)

type app_list struct{}

func (app *app_list) String() string {
	return "app_list"
}

func (app *app_list) Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger) {
	data.InitLogger(logger)
	apps := &data.App{}
	code := 1000
	ds := apps.List()
	if ds == nil {
		code = 4000
	}
	result, _ := format(code, ds)
	w.Write(result)
}

func init() {
	var app_handlers map[string]Handler = map[string]Handler{
		"app/list": &app_list{},
	}
	for p, h := range app_handlers {
		Services[p] = h
	}
}
