package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
)

type FastCGI struct{}

func (s *FastCGI) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	r.Write([]byte("hello, FastCGI"))
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8989")
	srv := new(FastCGI)
	fcgi.Serve(listener, srv)
	select {}
}
