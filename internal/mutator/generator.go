package mutator

func GenerateBody(fields map[string]string) map[string]interface{} {
	body := make(map[string]interface{})

	for field, typ := range fields {
		switch typ {

		case "string":
			body[field] = "test"

		case "integer":
			body[field] = 25

		case "boolean":
			body[field] = true

		default:
			body[field] = nil
		}
	}

	return body
}
