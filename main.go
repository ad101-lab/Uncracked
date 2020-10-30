// main.go

package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var indexTpl = template.Must(template.ParseFiles("html/index.html"))

func main()  {
	fmt.Println("Starting server on http://localhost:3000/");
	
	mux := http.NewServeMux()


	var fs = http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":3000", nil);
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, nil)
}
