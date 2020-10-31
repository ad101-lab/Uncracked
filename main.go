// main.go

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

type pageVariables struct {
	username string
	image    string
}

var indexTpl = template.Must(template.ParseFiles("assets/index.html"))

func main() {
	fmt.Println("Starting server on http://localhost:3000/")

	http.HandleFunc("/signed-up", signupHandler)

	http.ListenAndServe(":3000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Hash user info

		// Search user info in database

		// Verify user data

		// Send response

		// Forward to dashboard

	} else {
		// Reload page, with error
	}
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
