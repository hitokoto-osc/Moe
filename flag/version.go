package flag

import (
	"fmt"
	"github.com/hitokoto-osc/Moe/config"
	pflag "github.com/spf13/pflag"
	"os"
)

// V is a flag mapping, intended to show version text if be true
var V bool

func registerVersionFlag() {
	pflag.BoolVarP(&V, "version", "v", false, "查看版本信息")
}

func handleVersionFlag() {
	if V {
		fmt.Printf("Moe, A lightweight hitokoto status data merger. Authored by a632079\n版本: %s\n版本哈希: %s\n提交时间: %s\n编译时间: %s\n编译环境: %s\n", config.Version, config.BuildTag, config.CommitTime, config.BuildTime, config.GoVersion)
		os.Exit(0)
	}
}
