package database

import (
	"fmt"
	"go.uber.org/zap"

	// 导入 MySQL 库用于 sqlx 驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/jmoiron/sqlx"
)

// DB 是 sqlx 的示例
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

// InitDB 用于初始化数据库连接
func InitDB() {
	defer zap.L().Sync()
	zap.L().Info("[database] 正在与数据库建立连接...")
	var connectionURI = getConnectionURI
	zap.L().Debug("[database] 数据库连接地址：" + connectionURI())
	DB = sqlx.MustConnect("mysql", connectionURI())
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		zap.L().Fatal("连接数据库失败！", zap.Error(err))
	}
}
