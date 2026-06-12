package analyzer

type Finding struct {
	Method		string
	Path		string
	Status		int
	Payload		map[string]any
	Reason		string
}
