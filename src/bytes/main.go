package main

import "bytes"
// http://blog.csdn.net/luomoshusheng/article/details/50768750
func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))

	// 直接new初始化
	b1 := new(bytes.Buffer)
	b2 := bytes.NewBuffer([]byte("go"))
	b3 := bytes.NewBufferString("go")
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}