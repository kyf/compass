package data

import (
	"log"
)

type App struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Icon   string      `json:"icon"`
	Desc   string      `json:"desc"`
	Apk    string      `json:"apk"`
	Logger *log.Logger `json:"-"`
}

func (a *App) List() []App {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}
	sql := "select `id`, `name`, `icon`, `desc`, `apk` from `app` order by id desc limit 10"
	rows, err := db.Query(sql)
	if err != nil {
		a.Logger.Printf("app list err is %v", err)
		return nil
	}
	defer rows.Close()

	result := make([]App, 0)
	for rows.Next() {
		var id int
		var name, icon, desc, apk string
		err = rows.Scan(&id, &name, &icon, &desc, &apk)
		if err != nil {
			a.Logger.Printf("err is %v", err)
			continue
		}
		result = append(result, App{id, name, icon, desc, apk})
	}

	return result
}

func (a *App) Add(name, icon, desc, apk string) error {
	db, err := initDB()
	if err != nil {

	}
}

func (a *App) Update() {

}

func (a *App) Remove() {

}
