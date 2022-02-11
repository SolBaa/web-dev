package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	r := mux.NewRouter()
	if err != nil {
		fmt.Println(err)
	}
	r.HandleFunc("/", hello)
	http.ListenAndServe(":8080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Holaaa")
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		fmt.Println(err)
	}

}
