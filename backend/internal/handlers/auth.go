func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&req)

	var hashed string
	err := db.DB.QueryRow("SELECT password FROM users WHERE email=$1", req.Email).Scan(&hashed)
	if err != nil {
		http.Error(w, "unauthorized", 401)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(req.Password)) != nil {
		http.Error(w, "unauthorized", 401)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	json.NewEncoder(w).Encode(map[string]string{"token": t})
}