package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	http.ListenAndServe(":8080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Holaaa")
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	data := User{Name: "Sol"}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}

}
