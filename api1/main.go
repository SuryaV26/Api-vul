package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Flag     string `json:"flag,omitempty"`
	Message  string `json:"message,omitempty"`
}

var users = map[int]User{
	1: {ID: 1, Username: "Luffy"},
	2: {ID: 2, Username: "Sanji"},
	3: {ID: 3, Username: "Zoro"},
	4: {ID: 4, Username: "Gol D Roger"},
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Check for the bypass token in headers
	Token := r.Header.Get("X-Authorization")
	encodedID := base64.StdEncoding.EncodeToString([]byte(idStr))
	if id == 4 && Token != encodedID {
		user := User{
			Message: "It seems there's a protection. Try adding some header.",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(user)
		return
	}

	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Add the flag to the response if the bypass token is correct
	if id == 4 && Token == encodedID {
		user.Flag = "flag{API1-BOLA}"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/api/users", getUser)
	http.ListenAndServe(":8033", nil)
}
