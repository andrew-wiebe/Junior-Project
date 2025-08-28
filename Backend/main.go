package main

//Do "go run main.go" to start the backend
//net/http is the thing that allows for apis to work in go
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the API!")
	fmt.Fprintln(w, fmt.Sprintf("Request method: %s", r.Method))
}
