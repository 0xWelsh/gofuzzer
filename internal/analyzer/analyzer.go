package analyzer

import "fmt"

func Analyze(
	status int,
	method string,
	path string,
) {

	if status >= 500 {
		fmt.Printf(
			"[!] Possible issue: %s %s -> %d\n",
			method,
			path,
			status,
		)
	}
}
