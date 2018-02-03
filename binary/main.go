package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var pi float64
	bpi := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(bpi)
	err := binary.Read(buf, binary.LittleEndian, &pi)

	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Print(pi)

	var pi2 float64 = 3.141592653589793
	buf2 := new(bytes.Buffer)
	err = binary.Write(buf2, binary.LittleEndian, pi2)

	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Print(buf2.Bytes())
}
