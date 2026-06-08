package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:

			json.NewEncoder(w).Encode(
				map[string]string{
					"message": "list pets",
				},
			)

		case http.MethodPost:

			var body map[string]interface{}

			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

				http.Error(
					w,
					"invalid json",
					http.StatusBadRequest,
				)

				return
			}

			name, _ := body["name"].(string)

			if name == "../../../etc/passwd" {

				http.Error(
					w,
					"internal error",
					http.StatusInternalServerError,
				)

				return
			}

			json.NewEncoder(w).Encode(
				map[string]interface{}{
					"received": body,
				},
			)

		default:

			http.Error(
				w,
				"method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})

	log.Println("Listening on :8080")

	log.Fatal(
		http.ListenAndServe(
			":8080",
			nil,
		),
	)
}
