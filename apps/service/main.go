package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/kyf/compass/service"
)

const (
	SERVICE_ADDR string = ":9968"
)

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()

	for p, s := range service.Services {
		m.Get(fmt.Sprintf("/%s", p), s.Handle)
	}

	m.RunOnAddr(SERVICE_ADDR)
}
