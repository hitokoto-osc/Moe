package flag

import "github.com/spf13/pflag"

func init() {
	// Register Flag mapping
	registerVersionFlag()
	registerHelpFlag()
	registerDebugFlag()
	registerConfigFlag()

	// Parse Flag
	pflag.Parse()
}

// Do is a func will be called at init, registering the drivers of program
func Do() {
	handleVersionFlag()
	handleHelpFlag()
}
