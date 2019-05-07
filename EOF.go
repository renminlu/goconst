package goconst

import (
	"runtime"
)

var (
	EOL string
)

func init() {

	switch runtime.GOOS {
	case "windows":
		EOL = "\r\n"
	default:
		EOL = "\n"
	}
}
