package flag

import (
	"github.com/hitokoto-osc/Moe/config"
	"github.com/spf13/pflag"
)

func registerConfigFlag() {
	pflag.StringVarP(&config.File, "config", "c", "", "设定档案信息")
}
