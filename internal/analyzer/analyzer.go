package analyzer

import (
	"fmt"

	"gofuzzer/internal/ui"
)

func Analyze(
	status int,
	method string,
	path string,
) {

	if status >= 500 {
		fmt.Printf(
			"%s[!] Possible issue:%s %s %s -> %d\n",
			ui.Red,
			ui.Reset,
			method,
			path,
			status,
		)
	}
}
