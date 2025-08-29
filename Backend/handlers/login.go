package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	mockUserList := map[string]string{
		"user1": "password1",
		"user2": "password2",
	}

	fmt.Fprintln(w, "Hello from the login endpoint!")

	if r.Method == http.MethodPost {
		// Handle login logic here
		fmt.Fprintln(w, "Logging in...")
		username := r.FormValue("username")
		password := r.FormValue("password")

		if storedPassword, ok := mockUserList[username]; ok && storedPassword == password {
			fmt.Fprintln(w, "Login successful!")
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
