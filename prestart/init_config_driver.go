package prestart

import (
	"bytes"
	"github.com/hitokoto-osc/Moe/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

// The Config Parse Driver is served by viper
func initConfigDriver() {
	config.SetDefault()
	// Parse env config
	viper.SetEnvPrefix("moe") // like: MOE_PORT=8000
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default viper information
	viper.SetConfigName("config")
	viper.SetConfigType("toml") // Toml is the best!
	if config.File != "" {
		content, err := ioutil.ReadFile(config.File)
		if err != nil {
			log.Fatalf("[prestart] can't read specific config file, path: %s \nerror detail: %s\n", config.File, err)
		}
		err = viper.ReadConfig(bytes.NewBuffer(content))
		if err != nil {
			log.Fatalf("[prestart] can't load specific config file, path: %s \nerror detail: %s\n", config.File, err)
		}
	} else {
		// Parse path etc > home > localPath
		viper.AddConfigPath("/etc/.Moe")
		viper.AddConfigPath("$HOME/.Moe")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../conf")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("[prestart] Fatal error while reading config file: %s \n", err)
		}
	}
	config.Inject()
}
