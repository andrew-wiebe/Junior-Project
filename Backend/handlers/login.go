package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from the login endpoint!")
	// if r.Method == http.MethodPost {
	// 	// Handle login logic here
	// 	fmt.Fprintln(w, "Logging in...")
	// } else {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// }
}
