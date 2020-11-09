package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"NobodyNothin", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/Welcome.html"))

	http.HandleFunc("/", Listener)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func Listener(w http.ResponseWriter, r *http.Request) {

	if name := r.FormValue("name"); name != "" {
		welcome := name
	}

	if err := templates.ExecuteTemplate(w, "Welcome.html", Welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
