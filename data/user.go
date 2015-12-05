package data

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

type User struct {
	AdminName string      `json:"admin_name"`
	AdminPwd  string      `json:"admin_pwd"`
	Logger    *log.Logger `json:"-"`
}

func (u *User) Check(adminName, adminPwd string) bool {
	sql := "select count(1) as num from `admin` where admin_name = '%s' and admin_pwd = '%s'"
	db, err := initDB()
	if err != nil {
		u.Logger.Fatalf("initDB ERR:%v", err)
	}
	bin := md5.Sum([]byte(adminPwd))
	tmp := make([]byte, 32)
	hex.Encode(tmp, bin[:])
	rows, err := db.Query(fmt.Sprintf(sql, adminName, string(tmp)))
	if err != nil {
		u.Logger.Printf("user check err: sql is %s, err is %v", sql, err)
		return false
	}

	if rows.Next() {
		var num int
		rows.Scan(&num)
		if num > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
