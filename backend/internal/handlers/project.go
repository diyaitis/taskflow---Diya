package handlers

import (
	"encoding/json"
	"net/http"

	"taskflow/internal/db"
)

func Projects(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		rows, _ := db.DB.Query("SELECT id,name FROM projects")

		var res []map[string]string

		for rows.Next() {
			var id, name string
			rows.Scan(&id, &name)

			res = append(res, map[string]string{
				"id":   id,
				"name": name,
			})
		}

		json.NewEncoder(w).Encode(res)

	case "POST":
		var body struct {
			Name string `json:"name"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		_, _ = db.DB.Exec("INSERT INTO projects (id,name) VALUES (gen_random_uuid(),$1)", body.Name)

		json.NewEncoder(w).Encode(map[string]string{"message": "created"})
	}
}