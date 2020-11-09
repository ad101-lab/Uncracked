package main

import (
	"fmt"
	//"golang.org/x/crypto/bcrypt"
	"database/sql"
	"go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type DateVars struct {
	Date string
	Time string
}

type PageVars struct {
	UserName string
	Id       int
}

var IndexVisits = 0

const port = "80"

func main() {
	// Set handlers
	http.HandleFunc("/date", DateTime)
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Starting Server On: https://localhost:" + port)
	//Log the result
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func sqlLogin(Username, Password string) {

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	IndexVisits++
	w.Write([]byte(strconv.Itoa(IndexVisits)))
}

func DateTime(w http.ResponseWriter, r *http.Request) {

	// Create PageVars struct
	now := time.Now()
	HomePageVars := DateVars{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	// Open up the file in a template
	tpl, err := template.ParseFiles("datetime.html")

	// Log err
	if err != nil {
		log.Print("Template Parsing Error: ", err)
	}

	// Replace varables
	err = tpl.Execute(w, HomePageVars)

	// Log err
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
