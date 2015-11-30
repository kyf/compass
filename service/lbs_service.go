package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	ak          string = "C17e4ac75ae027235a4db6a79ccbe3e2"
	lbs_api_uri string = "http://api.map.baidu.com/location/ip?ak=%s&ip=%s&coor=bd09ll"
)

type lbsapi struct {
	Address string `json:"address"`
	Status  int    `json:"status"`
}

type lbs_is_china struct{}

func (lbs *lbs_is_china) String() string {
	return "lbs_is_china"
}

func (lbs *lbs_is_china) Handle(r *http.Request, w http.ResponseWriter, logger *log.Logger) {
	RemoteAddr := strings.Split(r.RemoteAddr, ":")
	ip := RemoteAddr[0]
	res, err := http.Get(fmt.Sprintf(lbs_api_uri, ak, ip))
	if err != nil {
		logger.Printf("invalid ip add [%s], err is %v", ip, err)
		result, _ := format(4000, nil)
		w.Write(result)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Printf("invalid ip add [%s], err is %v", ip, err)
		result, _ := format(4000, nil)
		w.Write(result)
		return
	}

	var jsonBody lbsapi
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		logger.Printf("invalid ip add [%s], err is %v", ip, err)
		result, _ := format(4000, nil)
		w.Write(result)
		return
	}

	if jsonBody.Status == 0 {
		var result []byte
		if strings.EqualFold(jsonBody.Address[:2], "CN") {
			result, _ = format(1000, nil)
		} else {
			result, _ = format(4000, nil)
		}

		w.Write(result)
		return
	} else {
		logger.Printf("invalid ip add [%s], err is %s", ip, string(body))
		result, _ := format(4000, nil)
		w.Write(result)
		return
	}
	result, _ := format(1000, nil)
	w.Write(result)
}

func init() {
	var lbs_handlers map[string]Handler = map[string]Handler{
		"lbs/ischina": &lbs_is_china{},
	}
	for p, h := range lbs_handlers {
		Services[p] = h
	}
}
