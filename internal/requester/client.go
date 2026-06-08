package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gofuzzer/internal/mutator"
	"gofuzzer/internal/parser"
)

func BuildJSONRequest(
	baseURL string,
	tc mutator.TestCase,
) (*http.Request, error) {

	bodyBytes, err := json.Marshal(tc.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		tc.Endpoint.Method,
		baseURL+tc.Endpoint.Path,
		bytes.NewBuffer(bodyBytes),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	return req, nil
}

func BuildRequest(baseURL string, ep parser.Endpoint) (*http.Request, error) {
	url := baseURL + ep.Path

	fmt.Printf("Building request: %s %s\n", ep.Method, url)

	return http.NewRequest(
		ep.Method,
		url,
		nil,
	)
}

func Send(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}
