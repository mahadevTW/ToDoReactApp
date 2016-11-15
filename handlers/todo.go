package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	"git.todo-app.com/ToDoReactApp/models"
	repo "git.todo-app.com/ToDoReactApp/repository"
)

type ToDoJSON struct {
	Item string
	Id   int
}

func MeAliveMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	byteContents := []byte("Alive")
	w.Write(byteContents)
}

func AddToDo(db *sql.DB, todoRepo repo.ToDoRepository) http.HandlerFunc {
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
		todoRepo.Insert(requestBody.Item, db)
		w.Write([]byte("Success"))
	}
}

func SelectToDos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, _ := models.ToDoSelectAll(db)
		todosJSON, err := json.Marshal(todos)
		if err != nil {
			fmt.Print("Error in json marshalling")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(todosJSON))
	}
}

func DeleteToDoHandler(db *sql.DB, todoRepo repo.ToDoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		requestBody := &ToDoJSON{}
		_ = json.Unmarshal(body, requestBody)
		fmt.Println("REQUEST BODY :", requestBody.Id)
		err := todoRepo.Delete(db, requestBody.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
