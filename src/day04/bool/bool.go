package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
	var b = false
	fmt.Println("b=",b)
	fmt.Println("b.type=",reflect.TypeOf(b))
	fmt.Println("b.size=",unsafe.Sizeof(b))
	//bool类不可以使用0和1来赋值或进行判断
}