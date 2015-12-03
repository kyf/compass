package data

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	AdminName string `json:"admin_name"`
	AdminPwd  string `json:"admin_pwd"`
}

func (u *User) Check(adminName, adminPwd string) bool {
	sql := "select count(1) as num from `admin` where admin_name = '%s' and admin_pwd = '%s'"
	db := initDB()
	bin := md5.Sum([]byte(adminPwd))
	tmp := make([]byte, 32)
	hex.Encode(tmp, bin[:])
	rows, err := db.Query(fmt.Sprintf(sql, adminName, string(tmp)))
	if err != nil {
		logger.Printf("Check err: sql is %s, err is %v", sql, err)
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
