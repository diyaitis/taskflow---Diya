package handlers

import (
	"encoding/json"
	"net/http"

	"taskflow/internal/db"
)

func Tasks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		rows, err := db.DB.Query("SELECT id, title FROM tasks")
		if err != nil {
			w.WriteHeader(500)
			return
		}

		var tasks []map[string]string

		for rows.Next() {
			var id, title string
			rows.Scan(&id, &title)

			tasks = append(tasks, map[string]string{
				"id":    id,
				"title": title,
			})
		}

		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var body struct {
			Title string `json:"title"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		_, err := db.DB.Exec(
			"INSERT INTO tasks (id, title) VALUES (gen_random_uuid(), $1)",
			body.Title,
		)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "task created",
		})
	}
}