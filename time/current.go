package main

import (
	//"time"
	"fmt"
	"time"
	//"strconv"
)

func main(){
	//t1:=time.Now().Year()        //年
	//
	//t2:=time.Now().Month()      //月
	//
	//t3:=time.Now().Day()         //日
	//
	//t4:=time.Now().Hour()        //小时
	//
	//t5:=time.Now().Minute()      //分钟
	//
	//t6:=time.Now().Second()      //秒
	//
	//var month int = int(t2)
	//fmt.Print(t1, month, t3, t4, t5, t6)
	//str := strconv.FormatInt(int64(t1), 10)
	//fmt.Print(str)
	// 44 46 31 38 30 38
	//hexByte, _:= hex.DecodeString("444631383038")
	//fmt.Printf("Binary: %16b", hexByte)
	current := time.Now()
	currentStr := current.Format("20060102150405")
	dateByte := []byte(currentStr[2:])
	fmt.Print(currentStr[2:])
	hexByte := make([]byte, 12)
	for i := 0; i < len(currentStr); i++{
		hexByte = append(hexByte, dateByte[i])
	}
	fmt.Print(dateByte)

}
