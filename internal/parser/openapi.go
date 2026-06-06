package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type OpenAPI struct {
	Paths map[string]map[string]interface{} `yaml:"paths"`
}


func Parse(specPath string) ([]Endpoint, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, err
	}

	var spec OpenAPI

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	var endpoints []Endpoint

	for path, methods := range spec.Paths {
		for method := range methods {
			endpoints = append(endpoints, Endpoint{
				Method: method,
				Path: path,
			})
		}
	}

	return endpoints, nil
}
