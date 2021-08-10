package main

import "fmt"

func main() {
	//交换赋值
	a := 2
	b := 3
	t := a
	a = b
	b = t
	fmt.Printf("After Change That A= %v B= %v\n", a, b)
	//符合型赋值
	a += 10
	fmt.Printf("After Added That A= %v\n", a)
	//基础表达式
	var d int
	d = a
	d = 8 + 2*8
	d = 890
	fmt.Println("d Values is=", d)
	//面试题：不使用中间变量完成变量值之间的交换
	var num1 int = 6
	var num2 int = 8
	num1 = num1 + num2
	num2 = num1 - num2 //此时num2 = num1 + num2 - num2 ==> num2 = num1
	num1 = num1 - num2 //此时num1 = num1 + num2 - num1 ==> num1 = num2
	fmt.Printf("After change that num1 = %v and num2 = %v ", num1, num2)
}
