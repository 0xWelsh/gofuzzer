package engine

import (
	"encoding/json"
	"fmt"

	"gofuzzer/internal/mutator"
	"gofuzzer/internal/config"
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
			"https://api.example.com",
			ep,
		)

		if err != nil {
			return err
		}

		fmt.Printf("%s %s\n", req.Method, req.URL.String())

		fmt.Printf("%s %s\n", ep.Method, ep.Path)

		for field, typ := range ep.Body {
			fmt.Printf(" %s (%s)\n", field, typ)

		}

		if len(ep.Body) > 0 {

			body := mutator.GenerateBody(ep.Body)

			jsonBody, err := json.MarshalIndent(
				body,
				"",
				" ",
			)

			if err != nil {
				return err
			}

			fmt.Println(string(jsonBody))
		}
	}

	return nil
}
