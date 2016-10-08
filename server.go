package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"git.todo-app.com/ToDoReactApp/handlers"
	repo "git.todo-app.com/ToDoReactApp/repository"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db, err := sql.Open("postgres", "user=todo_user password=todo dbname=todo_app sslmode=disable")
	defer db.Close()
	if err != nil {
	}
	todoRepo := &repo.ToDo{}
	AddToDoHandler := handlers.AddToDo(db)
	DeleteToDoHandler := handlers.DeleteToDoHandler(db, todoRepo)

	r.HandleFunc("/alive", handlers.MeAliveMethod)
	r.HandleFunc("/todos", handlers.SelectToDos(db)).Methods("GET")
	r.HandleFunc("/todo", AddToDoHandler).Methods("POST")
	r.HandleFunc("/todo", DeleteToDoHandler).Methods("DELET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./out/build/")))
	log.Println("Server started: http://localhost:" + port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
