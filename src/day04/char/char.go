package main

import "fmt"

func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	//这样直接输出会输出对应字符的ASCII码
	fmt.Println("c1=",c1)//=>97
	fmt.Println("c2=",c2)//=>46
	//如果我们希望是出对应字符时，则需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)
	//使用如下方法存储字符，传统byte型会出现溢出
	var c3 string = "北"
	fmt.Printf("c3=%s",c3)
	//报错案例:要将byte类型转换为int型
	var c4 int = '京'
	fmt.Printf("\nc4=%c c4的码值=%d",c4,c4)
	//使用数值反输出UTF-8字符
	var c5 int = 22269// 22269 -》 '国'
	fmt.Printf("\nc5=%c",c5)
	//字符类是可以进行运算的，相当于一个整数，运算时是按照码值运算
	var n1 = 10 + 'a'
	fmt.Println("\nn1=",n1)
	fmt.Printf("n1=%c",n1)
}