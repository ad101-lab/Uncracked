// main.go

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type signupData struct {
	Business string
	Email    string
	password string
	city     string
	state    int8
	zipCode  uint32
	addrLine string
}

var indexTpl = template.Must(template.ParseFiles("html/index.html"))

func main() {
	fmt.Println("Starting server on http://localhost:3000/")

	http.HandleFunc("/signed-up", signupHandler)

	http.ListenAndServe(":3000", nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Check user info

		// Hash password

		// Add user to database

		// Send us email
	} else {
		// Redirect to Sign up
	}
}
