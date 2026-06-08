package engine

import (
	"encoding/json"
	"fmt"

	"gofuzzer/internal/analyzer"
	"gofuzzer/internal/config"
	"gofuzzer/internal/mutator"
	"gofuzzer/internal/parser"
	"gofuzzer/internal/requester"
)

type Engine struct {
	cfg      config.Config
	specPath string
}

func New(cfg config.Config, specPath string) (*Engine, error) {
	return &Engine{
		cfg:      cfg,
		specPath: specPath,
	}, nil
}

func (e *Engine) Run() error {
	fmt.Printf("Loading spec: %s\n", e.specPath)
	fmt.Printf("Workers: %d\n", e.cfg.Workers)

	endpoints, err := parser.Parse(e.specPath)
	if err != nil {
		return err
	}

	fmt.Println("\nDiscovered requests.")

	for _, ep := range endpoints {
		req, err := requester.BuildRequest(
			"http://localhost:8080",
			ep,
		)

		resp, err := requester.Send(req)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
			continue
		}

		fmt.Printf(
			"[%d] %s %s\n",
			resp.StatusCode,
			req.Method,
			req.URL.String(),
		)

		resp.Body.Close()

		if err != nil {
			return err
		}

		fmt.Printf("%s %s\n", req.Method, req.URL.String())

		fmt.Printf("%s %s\n", ep.Method, ep.Path)

		for field, typ := range ep.Body {
			fmt.Printf(" %s (%s)\n", field, typ)

		}

		if len(ep.Body) > 0 {

			testCases := mutator.MutateBody(ep)

			for _, tc := range testCases {

				req, err := requester.BuildJSONRequest(
					"http://localhost:8080",
					tc,
				)

				resp, err := requester.Send(req)
				if err != nil {
					fmt.Printf("Request failed: %v\n", err)
					continue
				}


				fmt.Printf(
					"[%d] %s %s\n",
					resp.StatusCode,
					req.Method,
					req.URL.String(),
				)

				analyzer.Analyze(
					resp.StatusCode,
					req.Method,
					req.URL.Path,
				)

				resp.Body.Close()

				if err != nil {
					return err
				}

				fmt.Printf(
					"%s %s\n",
					req.Method,
					req.URL.String(),
				)

				jsonBody, err := json.MarshalIndent(
					tc.Body,
					"",
					" ",
				)

				if err != nil {
					return err
				}

				fmt.Println(string(jsonBody))
				fmt.Println()
			}
		}
	}

	return nil
}
