// +build ignore
package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

var templates = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request){
	task := r.FormValue("task")
	date := r.FormValue("date")
	deadline := r.FormValue("deadline")
	fmt.Println(task)
	fmt.Println(date)
	fmt.Println(deadline)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/save", save)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

