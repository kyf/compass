package data

import ()

type Setting struct {
	Id     int `json:"id"`
	AdShow int `json:"ad_show"`
}

func (s *Setting) Read() {
	db := initDB()
	sql := "select `id`, `ad_show` from `setting` limit 1"
	rows, err := db.Query(sql)
	if err != nil {
		logger.Printf("sql error: sql is %s, err is %v", sql, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		var id int
		var ad_show int
		err = rows.Scan(&id, &ad_show)
		if err != nil {
			logger.Printf("scan error: err is %v", err)
			return
		}
		s.Id = id
		s.AdShow = ad_show
	}
}

func (s *Setting) Write() {

}
