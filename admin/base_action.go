package admin

import (
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request, s sessions.Session, logger *log.Logger)

var (
	ActionHandlers map[string]handler = make(map[string]handler)
)
