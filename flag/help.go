package flag

import (
	"fmt"
	"os"

	"github.com/hitokoto-osc/Moe/config"
	pflag "github.com/spf13/pflag"
)

var h bool

func registerHelpFlag() {
	pflag.BoolVarP(&h, "help", "h", false, "查看程序帮助")
}

func handleHelpFlag() {
	if h {
		fmt.Printf(`Moe v%s
使用: moe [-dhv] [-c filename]
选项：
`, config.Version)
		pflag.PrintDefaults()
		os.Exit(0)
	}
}
