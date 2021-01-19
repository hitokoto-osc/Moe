package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func getConnectionURI () string {
	return fmt.Sprintf("user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True")
}

func InitDB () {
	DB = sqlx.MustConnect("mysql", getConnectionURI())
}
