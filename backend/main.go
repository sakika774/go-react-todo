package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var (
	todos []Todo
	idSeq = 1
	mu    sync.Mutex
	file  = "todos.json"
)

// ファイル読み込み
func loadTodos() {
	data, err := ioutil.ReadFile(file)
	if err == nil {
		json.Unmarshal(data, &todos)
		for _, t := range todos {
			if t.ID >= idSeq {
				idSeq = t.ID + 1
			}
		}
	}
}

// ファイル保存
func saveTodos() {
	data, _ := json.MarshalIndent(todos, "", "  ")
	ioutil.WriteFile(file, data, 0644)
}

func main() {
	loadTodos()

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
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
			saveTodos()
			mu.Unlock()
			json.NewEncoder(w).Encode(t)
			return
		}
		
		if r.Method == http.MethodDelete {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	mu.Lock()
	defer mu.Unlock()
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			break
		}
	}
	w.WriteHeader(http.StatusOK)
	return
		}
	})

	http.HandleFunc("/todos/toggle", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPatch {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var id int
		fmt.Sscanf(idStr, "%d", &id)

		mu.Lock()
		defer mu.Unlock()
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = !todos[i].Done
				saveTodos()
				break
			}
		}
		w.WriteHeader(http.StatusOK)
		return
	}
})
	
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}