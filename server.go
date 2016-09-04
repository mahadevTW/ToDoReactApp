package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"git.todo-app.com/ToDoReactApp/handlers"
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
	ToDoHandler := handlers.AddToDo(db)
	r.HandleFunc("/alive", handlers.MeAliveMethod)
	r.HandleFunc("/todos", handlers.SelectToDos(db)).Methods("GET")
	r.HandleFunc("/todo", ToDoHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./out/build/")))
	log.Println("Server started: http://localhost:" + port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
