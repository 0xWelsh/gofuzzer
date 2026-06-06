package engine

import (
	"fmt"

	"gofuzzer/internal/config"
	"gofuzzer/internal/parser"
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

	for _, ep := range endpoints {
		fmt.Printf("%s %s\n", ep.Method, ep.Path)
	}

	return nil
}
