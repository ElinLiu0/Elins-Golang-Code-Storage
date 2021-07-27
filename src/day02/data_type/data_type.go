package main

import "fmt"

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
	var c byte = 255
	fmt.Println("a=",a,"b=",b,"c=",c)
}
