package engine

import (
	"fmt"

	"gofuzzer/internal/config"
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

	return nil
}
