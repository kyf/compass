package data

import ()

type Setting struct {
	Id     int `json:"id"`
	AdShow int `json:"ad_show"`
}

func (s *Setting) Read() {
	sql := "select `id`, `ad_show` from `setting` limit 1"
	Query(sql)
}

func (s *Setting) Write() {

}
