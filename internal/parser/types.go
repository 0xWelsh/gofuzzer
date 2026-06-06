package parser

type Endpoint struct {
	Method string
	Path   string

	Body map[string]string
}
