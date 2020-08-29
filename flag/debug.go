package flag

import (
	"github.com/hitokoto-osc/Moe/config"
	pflag "github.com/spf13/pflag"
)

func registerDebugFlag() {
	pflag.BoolVarP(&config.Debug, "debug", "d", false, "启动调试模式")
}
