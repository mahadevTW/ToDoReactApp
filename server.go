package main


import(
	"net/http"
	"log"
	"os"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.Handle("/", http.FileServer(http.Dir("./out/build")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}