package service

import (
	"log"
	"net/http"
)

type lbs_is_china struct{}

func (lbs *lbs_is_china) String() string {
	return "lbs_is_china"
}

func (lbs *lbs_is_china) Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger) {
	w.Write([]byte("asdasdad"))
}

func init() {
	var lbs_handlers map[string]Handler = map[string]Handler{
		"lbs/ischina": &lbs_is_china{},
	}
	for p, h := range lbs_handlers {
		Services[p] = h
	}
}
