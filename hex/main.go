package main

import (
	"encoding/hex"
	"fmt"
	"crypto/md5"
)

func main(){
	//a := "30313233343536373238"
	address := "0123456728"
	by := []byte(address)
	//fmt.Print(by)
	h := md5.New()
	str := hex.EncodeToString(h.Sum(by))
	fmt.Print(str)
	//bs, err := hex.DecodeString(a)
	//src := make([]byte, 0)
	//hex.Encode(by, src)
	//fmt.Print(src)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(bs))
	var length int64 = 16
	fmt.Printf("%03X", length)
}
