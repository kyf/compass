package service

import (
	"log"
	"net/http"
)

type app_list struct{}

func (app *app_list) String() string {
	return "app_list"
}

func (app *app_list) Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger) {
	w.Write([]byte("asdasdad"))
}

func init() {
	var app_handlers map[string]Handler = map[string]Handler{
		"app/list": &app_list{},
	}
	for p, h := range app_handlers {
		Services[p] = h
	}
}
