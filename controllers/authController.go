package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/fariedrisky/go-restful-mysql/database"
	"github.com/fariedrisky/go-restful-mysql/models"
	"github.com/fariedrisky/go-restful-mysql/utils"
)

// Register registers a new user into the database.
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

// Login authenticates a user.
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	var storedUser models.User
	err := database.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", creds.Username).Scan(&storedUser.ID, &storedUser.Password)
	if err == sql.ErrNoRows || !utils.CheckPasswordHash(creds.Password, storedUser.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
