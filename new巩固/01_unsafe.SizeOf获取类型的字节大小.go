package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var(
		x byte
		x1 int
		x2 int32
		x3 int64
		x4 uint32
		x5 uint64
	)
	fmt.Println(x,x1,x2,x3,x4,x5)
	fmt.Println(unsafe.Sizeof(x1))  //unsafe.Sizeof 取值范围 一个字节是1-255

	var a int8 = 88
	var b uint8 = 10
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))


	str := "hello\"world\a"
	fmt.Println(str)

}
