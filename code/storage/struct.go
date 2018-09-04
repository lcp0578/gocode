package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
对齐， CPU更好地访问位于2字节（“2字节边界”）的倍数的地址处的2个字节，
并访问位于4字节边界上的4个字节，直到CPU的自然整数大小，
在现代CPU上是8字节（64位）。
**/
type MyData struct {
	aByte byte //1
	//anotherByte byte   //1
	aShort  int16  //2
	anInt32 int32  //4
	aSlice  []byte //24
}

func main() {
	dataType := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long \n", dataType.Size())
	n := dataType.NumField()
	for i := 0; i < n; i++ {
		field := dataType.Field(i)
		fmt.Printf("%s at offset %v,  size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	mdata := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}
	dataBytes := (*[32]byte)(unsafe.Pointer(&mdata))
	//因为大多数现代CPU都是Little-Endian：该值的最低位字节首先出现在内存中。
	fmt.Printf("Bytes are %#v\n", dataBytes)

	/**
	  type SliceHeader struct {
	      Data uintptr
	      Len  int
	      Cap  int
	  }
	  **/
	sliceBytes := *(*reflect.SliceHeader)(unsafe.Pointer(&mdata.aSlice))
	fmt.Printf("slice data is %#v\n", (*[3]byte)(unsafe.Pointer(sliceBytes.Data)))
}
