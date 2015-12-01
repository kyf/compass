package service

import (
	dal "github.com/kyf/compass/data"
	"log"
	"net/http"
)

type ad_is_show struct{}

func (a *ad_is_show) String() string {
	return "ad_is_show"
}

func (a *ad_is_show) Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger) {
	code := 1000
	setting := &dal.Setting{}
	setting.Read()

	echostr, err := format(code, nil)
	if err != nil {
		logger.Printf("err is %v", err)
		return
	}
	w.Write(echostr)
}

func init() {
	var ad_handlers map[string]Handler = map[string]Handler{
		"ad/show": &ad_is_show{},
	}
	for p, h := range ad_handlers {
		Services[p] = h
	}
}
