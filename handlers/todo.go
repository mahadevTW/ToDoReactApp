package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.todo-app.com/ToDoReactApp/models"
)

type ToDoJSON struct {
	Item string `json:"Item"`
}

func MeAliveMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	byteContents := []byte("Alive")
	w.Write(byteContents)
}

func AddToDo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Empty Body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		requestBody := &ToDoJSON{}
		err = json.Unmarshal(body, requestBody)
		if err != nil {
			return
		}
		models.ToDoInsert(requestBody.Item, db)
		w.Write([]byte("Success"))
	}
}
