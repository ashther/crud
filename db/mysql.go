package db

import (
	. "../utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Con is database connection
var Con *sql.DB

func init() {
	var err error
	Con, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/dsp")
	CheckErr(err)

	Con.SetMaxIdleConns(20)
	Con.SetMaxOpenConns(20)
	err = Con.Ping()
	CheckErr(err)

	_, err = Con.Exec("set names utf8;")
	CheckErr(err)
}
