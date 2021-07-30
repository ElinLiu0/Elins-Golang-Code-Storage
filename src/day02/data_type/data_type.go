package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//int8的最小范围是-128~127如果超出边界则会报错
	var num int8 = -128
	var num2 int8 = 127
	fmt.Println(num,num2)
	var num3 uint8 = 0
	//uint8最大范围到255
	var num4 uint8 = 255
	fmt.Println(num3,num4)
	// int , unit , rune,byte的使用
	var a int = 8000
	var b uint = 1
	//byte的极限值为255
	var c byte = 255
	fmt.Println("a=",a,"b=",b,"c=",c)
	//整形使用细节
	var n1 = 100//这里的n1是什么类型
	//%T代表显示后面指定对象的类型
	//fmt.Printf()可以用于做格式化输出
	fmt.Printf("n1的类型为：%T\n",n1)
	//查看变量的占用空间小的数据类型
	var n2 int64 = 10
	fmt.Printf("n2的数据类型为：%T，占用字节为：%d",n2,unsafe.Sizeof(n2))
	//保小不保大
	var age byte = 20
	fmt.Printf("\n年龄为：%d",age)
}
