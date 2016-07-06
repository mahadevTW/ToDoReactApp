package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"git.todo-app.com/ToDoReactApp/handlers"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.HandleFunc("/alive", handlers.MeAliveMethod)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./out/build/")))
	log.Println("Server started: http://localhost:" + port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
