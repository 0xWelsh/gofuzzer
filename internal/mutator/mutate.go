package mutator

import "gofuzzer/internal/parser"

func MutateBody(ep parser.Endpoint) []TestCase {

	var testCases []TestCase

	for field, typ := range ep.Body {

		if typ != "string" {
			continue
		}

		for _, payload := range StringPayloads {

			body := GenerateBody(ep.Body)

			body[field] = payload

			testCases = append(testCases, TestCase{
				Endpoint: ep,
				Body:     body,
			})
		}
	}

	return testCases
}
