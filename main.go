package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "Index", nil)
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
