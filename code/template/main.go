package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl.html")
	t, _ = template.ParseGlob("*.html")
	t = template.Must(template.ParseFiles("tmpl.html"))
	t.Execute(w, "hello word")
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
