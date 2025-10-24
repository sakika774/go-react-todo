package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		todos := []Todo{
			{ID: 1, Task: "React勉強"},
			{ID: 2, Task: "Go API作成"},
		}
		json.NewEncoder(w).Encode(todos)
	})

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}