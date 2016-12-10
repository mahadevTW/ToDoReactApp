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
	"github.com/gorilla/csrf"
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
	AddToDoHandler := handlers.AddToDo(db, todoRepo)
	DeleteToDoHandler := handlers.DeleteToDoHandler(db, todoRepo)

	CSRFProtector := csrf.Protect([]byte("32-byte-long-auth-key"))


	r.HandleFunc("/alive", handlers.MeAliveMethod)
	r.HandleFunc("/todos", handlers.SelectToDos(db, todoRepo)).Methods("GET")
	r.HandleFunc("/todo", AddToDoHandler).Methods("POST")
	r.HandleFunc("/todo", DeleteToDoHandler).Methods("DELETE")
	r.HandleFunc("/csrfToken", handlers.CSRFHandler()).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./out/build/")))
	log.Println("Server started: http://localhost:" + port)
	http.Handle("/", CSRFProtector(r))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
