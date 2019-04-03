package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/devices", devices)
	http.ListenAndServe("8080", nil)
}

func devices(wr http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq)
	tpl.Execute(wr, "index.gohtml")
}
