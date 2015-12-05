package data

import (
	"fmt"
	"log"
)

type Setting struct {
	Id     int         `json:"id"`
	AdShow int         `json:"ad_show"`
	Logger *log.Logger `json:"-"`
}

func (s *Setting) Read() {
	db, err := initDB()
	if err != nil {
		s.Logger.Fatalf("initDB ERR:%v", err)
	}
	sql := "select `id`, `ad_show` from `setting` limit 1"
	rows, err := db.Query(sql)
	if err != nil {
		s.Logger.Printf("setting read error: sql is %s, err is %v", sql, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		var id int
		var ad_show int
		err = rows.Scan(&id, &ad_show)
		if err != nil {
			s.Logger.Printf("scan error: err is %v", err)
			return
		}
		s.Id = id
		s.AdShow = ad_show
	}
}

func (s *Setting) Write(state int) {
	db, err := initDB()
	if err != nil {
		s.Logger.Fatalf("initDB ERR:%v", err)
	}
	sql := "update `setting` set `ad_show` = %d "
	s.Read()
	if s.Id == 0 {
		sql = "insert into `setting`(`ad_show`)values(%d)"
	}
	_, err = db.Exec(fmt.Sprintf(sql, state))
	s.Logger.Printf("setting write error: sql is %s, err is %v", sql, err)
}
