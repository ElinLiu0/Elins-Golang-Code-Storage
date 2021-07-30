package main

import "fmt"

func main() {
	var price float64
	price = 13.5
	fmt.Printf("当前价格为：%f\n", price*2.0)
	//单精度和双精度的差别
	var num1 float32 = -132.0000901
	var num2 float64 = -132.0000901
	fmt.Println("num1=",num1,"num2=",num2)
	//num1= -132.00009 num2= -132.0000901
	//十进制数形式：如：5.12，0.512
	num3 := 5.12
	num4 := .123
	fmt.Println("num3=",num3,"num4=",num4)
	//科学计数法
	num5 := 5.1234E2
	fmt.Println("num5=",num5)
	num6 := 5.1234E-2
	fmt.Println("num6=",num6)
}