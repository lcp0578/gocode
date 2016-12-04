package main

import (
	//	"fmt"
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = "/usr/local/go/bin/go"
		script := "/home/root/cgi-script" + r.URL.Path
		handler.Dir = "/home/root/cgi-script"
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8081", nil)
	select {} // 阻塞进程
}
