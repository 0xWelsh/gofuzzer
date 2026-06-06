package parser

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type OpenAPI struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

type Operation struct {
	RequestBody RequestBody `yaml:"requestBody"`
}

type RequestBody struct {
	Content map[string]MediaType `yaml:"content"`
}

type MediaType struct {
	Schema Schema `yaml:"schema"`
}

type Schema struct {
	Properties map[string]Property `yaml:"properties"`
}

type Property struct {
	Type string `yaml:"type"`
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

		for method, operation := range methods {

			body := make(map[string]string)

			if jsonMedia, ok := operation.RequestBody.Content["application/json"]; ok {

				for field, property := range jsonMedia.Schema.Properties {

					body[field] = property.Type
				}
			}

			endpoints = append(endpoints, Endpoint{
				Method: strings.ToUpper(method),
				Path: path,
				Body: body,
			})
		}
	}

	return endpoints, nil
}
