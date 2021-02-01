package config

import "github.com/spf13/viper"

// Inject 定义了 config 的注入入口（启动时改变 config）
func Inject() {
	DBHost = viper.GetString("database.host")
	DBPort = viper.GetInt("database.port")
	DBName = viper.GetString("database.dbname")
	DBUser = viper.GetString("database.user")
	DBPassword = viper.GetString("database.password")
}
