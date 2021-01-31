package config

import "github.com/spf13/viper"

func Inject() {
	DBHost = viper.GetString("database.host")
	DBPort = viper.GetInt("database.port")
	DBName = viper.GetString("database.dbname")
	DBUser = viper.GetString("database.user")
	DBPassword = viper.GetString("database.password")
}
