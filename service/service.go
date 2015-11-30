package service

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler interface {
	String() string
	Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger)
}

var (
	Services map[string]Handler = make(map[string]Handler)
	CodeMsg  map[int]string     = map[int]string{
		1000: "success",
		4000: "failure",
		4001: "server error",
	}
)

func format(code int, data interface{}) ([]byte, error) {
	result := map[string]interface{}{
		"code": code,
		"msg":  CodeMsg[code],
		"data": data,
	}

	if data == nil {
		delete(result, "data")
	}
	strresult, err := json.Marshal(result)
	return strresult, err
}
