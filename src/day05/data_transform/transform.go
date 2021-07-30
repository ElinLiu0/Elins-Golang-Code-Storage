package main

import (
	"fmt"
	"reflect"
	//在包名前添加_可以忽略导入
)

func main()  {
	var num int = 100
	//将其转换为float
	var changed float32 = float32(num)
	fmt.Println("before:",num,reflect.TypeOf(num))
	fmt.Println("after:",changed,reflect.TypeOf(changed))
	//低精度像高精度转换
	var low int8 = 24
	var high int16 = int16(low)
	fmt.Println("before:",low,reflect.TypeOf(low))
	fmt.Println("after:",high,reflect.TypeOf(high))
	fmt.Println("origin:",reflect.TypeOf(low))
	var num1 int64 = 39232
	var num2 int8 = int8(num1)
	//此时num2已经溢出int8最大范围，显示64
	fmt.Println("num2",num2)
}
