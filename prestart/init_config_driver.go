package prestart

import (
	"bytes"
	"github.com/cockroachdb/errors"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/hitokoto-osc/Moe/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

// The Config Parse Driver is served by viper
func initConfigDriver() {
	logger := logging.GetLogger()
	defer logger.Sync()
	config.SetDefault()

	// Parse env config
	viper.SetEnvPrefix("moe") // like: MOE_PORT=8000
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Set default viper information
	viper.SetConfigName("config")
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
		viper.AddConfigPath("./data") // docker purpose
		viper.AddConfigPath("../conf")
		viper.AddConfigPath("../config")
		err := viper.ReadInConfig()
		if err != nil {
			var e viper.ConfigFileNotFoundError
			if !errors.As(err, &e) {
				logger.Fatal("[init] Fatal error while reading config file.", zap.Error(err))
			}
			logger.Warn("[init] No config file detected, reading config from env.")
		}
	}

	logger.Debug("[init] config is parsed.",
		zap.String("config_file_used", viper.ConfigFileUsed()),
		zap.Any("settings", viper.AllSettings()),
	)
	config.Inject()
}
