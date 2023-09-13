package prestart

import (
	"bytes"
	"github.com/hitokoto-osc/Moe/logging"
	"go.uber.org/zap"
	"os"
	"strings"

	"github.com/hitokoto-osc/Moe/config"
	"github.com/spf13/viper"
)

// The Config Parse Driver is served by viper
func initConfigDriver() {
	logger := logging.GetLogger()
	defer logger.Sync()
	config.SetDefault()
	// Parse env config
	viper.SetEnvPrefix("moe") // like: MOE_PORT=8000
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default viper information
	viper.SetConfigName("config")
	viper.SetConfigType("toml") // Toml is the best!
	if config.File != "" {
		content, err := os.ReadFile(config.File)
		if err != nil {
			logger.Fatal(
				"[init] can't read specific config file.",
				zap.String("path", config.File),
				zap.Error(err),
			)
		}
		err = viper.ReadConfig(bytes.NewBuffer(content))
		if err != nil {
			logger.Fatal(
				"[init] can't load specific config file.",
				zap.String("path", config.File),
				zap.Error(err),
			)
		}
	} else {
		// Parse path etc > home > localPath
		viper.AddConfigPath("/etc/.Moe")
		viper.AddConfigPath("$HOME/.Moe")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./conf")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("../conf")
		viper.AddConfigPath("../config")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal("[init] Fatal error while reading config file.", zap.Error(err))
		}
	}
	config.Inject()
}
