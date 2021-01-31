package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var DB *sqlx.DB

func getConnectionURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}

func InitDB() {
	log.Info("[database] 正在与数据库建立连接...")
	var connectionURI = getConnectionURI
	log.Debug("[database]" + connectionURI())
	DB = sqlx.MustConnect("mysql", connectionURI())
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
