package main

import (
	"net/http"

	"text/template"

	"./cpu"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cpuload := cpu.GetCPULoadInfo()
	t := template.New("index template")

	t, _ = template.ParseFiles("template/index.html")

	t.Execute(w, cpuload)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
