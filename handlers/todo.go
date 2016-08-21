package handlers

import (
	"database/sql"
	"net/http"

	"git.todo-app.com/ToDoReactApp/models"
)

func MeAliveMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	byteContents := []byte("Alive")
	w.Write(byteContents)
}

func AddToDo(value string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		models.ToDoInsert(value, db)
		w.Write([]byte("Success"))
	}
}
