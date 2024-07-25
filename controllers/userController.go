package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fariedrisky/go-restful-mysql/database"
	"github.com/fariedrisky/go-restful-mysql/models"
)

// GetAllUsers fetches all users from the database.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
