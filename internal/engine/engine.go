package engine

import (
	"encoding/json"
	"fmt"
	"io"

	"gofuzzer/internal/analyzer"
	"gofuzzer/internal/config"
	"gofuzzer/internal/mutator"
	"gofuzzer/internal/parser"
	"gofuzzer/internal/requester"
	"gofuzzer/internal/ui"
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

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		length := len(bodyBytes)

		color := ui.Green

		if resp.StatusCode >= 400 {
			color = ui.Yellow
		}

		if resp.StatusCode >= 500 {
			color = ui.Red
		}

		fmt.Printf(
			"%s[%d]%s len=%d %s %s\n",
			color,
			resp.StatusCode,
			ui.Reset,
			length,
			req.Method,
			req.URL.String(),
		)

		resp.Body.Close()

		if err != nil {
			return err
		}

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

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return err
				}

				length := len(bodyBytes)

				color := ui.Green

				if resp.StatusCode >= 400 {
					color = ui.Yellow
				}

				if resp.StatusCode >= 500 {
					color = ui.Red
				}

				fmt.Printf(
					"%s[%d]%s len=%d %s %s\n",
					color,
					resp.StatusCode,
					ui.Reset,
					length,
					req.Method,
					req.URL.String(),
				)

				analyzer.Analyze(
					resp.StatusCode,
					req.Method,
					req.URL.Path,
				)

				resp.Body.Close()

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
