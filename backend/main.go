package main

import (
	"encoding/json"
	"net/http"
	"social-network/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := models.CreateTables()
	defer db.Close()
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/sayHi", sayHi)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello there"})
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hi"})
}
