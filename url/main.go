package main

import (
	"fmt"
	"net/url"
)

func main() {
	var uri = "http://www.baidu.com?a=33&b=2018年2月1日"
	encodeUri := url.QueryEscape(uri)
	fmt.Println(encodeUri)
	decodeUri, err := url.QueryUnescape(encodeUri)
	fmt.Println(err)
	fmt.Println(decodeUri)
}
