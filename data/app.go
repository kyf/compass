package data

import (
	"fmt"
	"log"
	"strings"
)

type App struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Icon   string      `json:"icon"`
	Desc   string      `json:"desc"`
	Apk    string      `json:"apk"`
	Logger *log.Logger `json:"-"`
}

func (a *App) Detail(id string) *App {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}
	sql := "select `id`, `name`, `icon`, `desc`, `apk` from `app` where `id` = %s"
	rows, err := db.Query(fmt.Sprintf(sql, id))
	if err != nil {
		a.Logger.Printf("app list err is %v", err)
		return nil
	}
	defer rows.Close()

	var result *App
	if rows.Next() {
		var id int
		var name, icon, desc, apk string
		err = rows.Scan(&id, &name, &icon, &desc, &apk)
		if err != nil {
			a.Logger.Printf("err is %v", err)
		} else {
			result = &App{id, name, icon, desc, apk, nil}
		}
	}

	return result
}

func (a *App) List() []App {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}
	sql := "select `id`, `name`, `icon`, `desc`, `apk` from `app` order by id desc limit 50"
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
		result = append(result, App{id, name, icon, desc, apk, nil})
	}

	return result
}

func (a *App) Add(name, icon, desc, apk string) error {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}

	sql := "insert into `app`(`name`, `icon`, `desc`, `apk`)values('%s', '%s', '%s', '%s')"
	_, err = db.Exec(fmt.Sprintf(sql, name, icon, desc, apk))
	if err != nil {
		a.Logger.Printf("sql err: sql is %s, err is %v", sql, err)
		return err
	}
	return nil
}

func (a *App) Update(id int, name, icon, desc, apk string) error {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}

	sql := "update `app` set `name` = '%s', `desc` = '%s'"

	if !strings.EqualFold(icon, "") {
		sql = strings.Join([]string{sql, fmt.Sprintf(" , `icon` = '%s'", icon)}, "")
	}

	if !strings.EqualFold(apk, "") {
		sql = strings.Join([]string{sql, fmt.Sprintf(" , `apk` = '%s'", apk)}, "")
	}

	sql = strings.Join([]string{sql, " where `id` = %d "}, "")

	_, err = db.Exec(fmt.Sprintf(sql, name, desc, id))
	if err != nil {
		a.Logger.Printf("app update err:sql is %s, err is %v", sql, err)
		return err
	}
	return nil
}

func (a *App) Remove(id int) error {
	db, err := initDB()
	if err != nil {
		a.Logger.Fatalf("initDB ERR:%v", err)
	}

	sql := "delete from `app` where `id` = %d"
	_, err = db.Exec(fmt.Sprintf(sql, id))
	if err != nil {
		a.Logger.Printf("app remove err: sql is %s, err is %v", sql, err)
		return err
	}

	return nil
}
