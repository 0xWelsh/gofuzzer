package main

import (
	"fmt"
	"log"
	"os"

	"gofuzzer/internal/config"
	"gofuzzer/internal/engine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gofuzzer <openapi-specs>")
		os.Exit(1)
	}

	specPath := os.Args[1]

	cfg := config.Default()

	fuzzer, err := engine.New(cfg, specPath)
	if err != nil {
		log.Fatalf("failed to initialize fuzzer: %v", err)
	}

	if err := fuzzer.Run(); err != nil {
		log.Fatalf("fuzzer error: %v", err)
	}
}
