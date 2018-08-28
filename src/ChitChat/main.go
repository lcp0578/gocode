package main

import (
	"net/http"
	"html/template"
)

func main(){
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("chit chat"))
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads();
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
