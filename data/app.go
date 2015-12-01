package data

import ()

type App struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
	Apk  string `json:"apk"`
}

func (a *App) List() []App {
	db := initDB()
	sql := "select `id`, `name`, `icon`, `desc`, `apk` from `app` order by id desc limit 10"
	rows, err := db.Query(sql)
	if err != nil {
		logger.Printf("app list err is %v", err)
		return nil
	}
	defer rows.Close()

	result := make([]App, 0)
	for rows.Next() {
		var id int
		var name, icon, desc, apk string
		err = rows.Scan(&id, &name, &icon, &desc, &apk)
		if err != nil {
			logger.Printf("err is %v", err)
			continue
		}
		result = append(result, App{id, name, icon, desc, apk})
	}

	return result
}

func (a *App) Add() {

}

func (a *App) Update() {

}

func (a *App) Remove() {

}
