package mutator

import "gofuzzer/internal/parser"

type TestCase struct {
	Endpoint parser.Endpoint
	Body     map[string]interface{}
}
