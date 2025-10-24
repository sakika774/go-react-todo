package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var (
	todos []Todo
	idSeq = 1
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodGet {
			mu.Lock()
			defer mu.Unlock()
			json.NewEncoder(w).Encode(todos)
			return
		}

		if r.Method == http.MethodPost {
			var t Todo
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			mu.Lock()
			t.ID = idSeq
			idSeq++
			todos = append(todos, t)
			mu.Unlock()
			json.NewEncoder(w).Encode(t)
			return
		}
	})

	// 初期データ
	todos = append(todos, Todo{ID: idSeq, Task: "React勉強"})
	idSeq++
	todos = append(todos, Todo{ID: idSeq, Task: "Go API作成"})
	idSeq++

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}