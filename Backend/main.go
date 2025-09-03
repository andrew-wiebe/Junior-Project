package main

//Do "go run main.go" to start the backend
//net/http is the thing that allows for apis to work in go
import (
	"fmt"
	"net/http"

	//This is how you import from within your own project
	handlers "github.com/andrew-wiebe/Backend/Handlers"
)

// http://localhost:8080/api
func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}

// We will use this type of thing to wrap around our api calls and allow them to branch off
// Make sure to nest them properly
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the API!")
	http.HandleFunc("/api/login", handlers.LoginHandler)


}
