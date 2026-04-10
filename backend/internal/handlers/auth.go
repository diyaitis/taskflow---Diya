package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"taskflow/internal/db"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	_, err := db.DB.Exec(
		"INSERT INTO users (id,email,password) VALUES (gen_random_uuid(),$1,$2)",
		req.Email,
		string(hashed),
	)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"user exists"}`))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "user created"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	var hashed string
	err := db.DB.QueryRow("SELECT password FROM users WHERE email=$1", req.Email).Scan(&hashed)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(`{"error":"unauthorized"}`))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(req.Password))
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(`{"error":"unauthorized"}`))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	json.NewEncoder(w).Encode(map[string]string{"token": t})
}