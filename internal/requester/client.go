package requester

import (
	"fmt"
	"net/http"

	"gofuzzer/internal/parser"
)

func BuildRequest(baseURL string, ep parser.Endpoint) (*http.Request, error) {
	url := baseURL + ep.Path

	fmt.Printf("Building request: %s %s\n", ep.Method, url)

	return http.NewRequest(
		ep.Method,
		url,
		nil,
	)
}
