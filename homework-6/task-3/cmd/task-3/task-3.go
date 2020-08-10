package main

import (	
	"fmt"
	"html/template"
	"net/http"
)

// Our page template
var pageTemplate *template.Template

// handleRequest handles a request like /hello?name=value
func handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// get 'name' parametre
	var name string
	if len(r.Form["name"]) != 0 {
		name = r.Form["name"][0]
	}

	// prepare a data for the page template
	data := struct { Name string }{ 
		Name: name,
	}
	
	// apply the data with the page template
	pageTemplate.Execute(w, data)
}

func main() {
	// load the page template
	pageTemplate = template.Must(template.ParseFiles("./html/page.html"))

	// initialize and start http server
	http.HandleFunc("/hello", handleRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is unable to start")
	} else {
		fmt.Println("Server has been successfully started")
	}
}
