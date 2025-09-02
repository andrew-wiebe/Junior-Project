package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: Put a database call here

	fmt.Fprintln(w, "Hello from the login endpoint!")

	if r.Method == http.MethodPost {
		// Handle login logic here
		if dbLoginCheck(r.FormValue("username"), r.FormValue("password")) {
			fmt.Fprintln(w, "Login successful!")
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func dbLoginCheck(username string, password string) bool {
	// TODO: Implement actual database call

	mockUserList := map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
	return mockUserList[username] == password
}
