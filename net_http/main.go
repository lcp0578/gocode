package main

import (
	//"fmt"
	//"io/ioutil"
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("hello"))
	io.WriteString(w, "Hello http server")
}
func main() {
	//	response, _ := http.Get("http://baidu.com")
	//	defer response.Body.Close()
	//	body, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(body))
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8081", nil)
}
