package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

type Dog struct {
	Breed    string
	WeightKg int
}

type Page struct {
	Json string
	Go string
}

func Index(w http.ResponseWriter, r *http.Request) {
	d := Dog{
		Breed:    "dalmation",
		WeightKg: 45,
	}
	b, _ := json.Marshal(d)
	json := Page{Json: string(b)}
    tmpl.ExecuteTemplate(w, "Index", json)
}

func main () {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)

	http.ListenAndServe(":8080", nil)
}

// type Dog struct {
// 	Breed    string
// 	WeightKg int
// }

// func main() {
// 	d := Dog{
// 		Breed:    "dalmation",
// 		WeightKg: 45,
// 	}
// 	b, _ := json.Marshal(d)
// 	fmt.Println(string(b))
// }
