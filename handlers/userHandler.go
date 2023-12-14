package handlers

import (
	"encoding/json"
	"net/http"

	"latihan1/models"
)

// simpan data pengguna
var users []models.User

// func get.. create...
func GetUser(w http.ResponseWriter, t *http.Request){
	user := models.User{ID:1, Name:"John Doe", Email: "john@example.com"}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser models.User

	//decode request body ke struct newUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tambahkan newUser ke slice Users
	users = append(users, newUser)

	// konfigurasi header dan response nya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}


