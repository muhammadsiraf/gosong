package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apa khabar")
}

func salam(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "Faricz",
		"Message": "Are you watching closely",
	}

	var t, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)
	http.HandleFunc("/salam", salam)
	fmt.Println("starting web server at http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
