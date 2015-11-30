package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST     string = "127.0.0.1"
	DB_PORT     string = "3306"
	DB_PROTOCOL string = "tcp"
	DB_USER     string = "root"
	DB_PWD      string = "123456"
	DB_NAME     string = "compass"
)

func initDB() {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", DB_USER, DB_PWD, DB_PROTOCOL, fmt.Sprintf("%s:%s", DB_HOST, DB_PORT), DB_NAME)
	dbHandler, err := sql.Open(dsn)
	if err != nil {
		logger.Fatalf("init db failure! err is %v", err)
	}

}
