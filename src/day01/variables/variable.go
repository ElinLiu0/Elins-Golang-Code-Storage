package main

import "fmt"

func getVal(num1 int,num2 int) (int,int){
	sum := num1 + num2
	sub := num1 - num2
	return sum ,sub
}
func main()  {
	sum, sub := getVal(4,5)
	fmt.Println("sub=",sub,"sum=",sum)
	sum2 ,_ := getVal(10,30)
	fmt.Println("sum",sum2)
	//第一种：指定类型后将其赋值
	//声明一个整数型的变量i
	var i int
	//在不被赋值时i为0，i=10使得在内存中将其赋值为10
	i = 10
	//打印
	fmt.Println("An i has been assign in",i)
	//第二种：直接赋值，由GoLang编译器自己推倒数据类型
	var num = 1
	fmt.Println("num=",num)
	//第三种：省略var,注意 := 左侧的变量应该是已经声明过的，否则会导致编译错误
	//下面的方式等价于 var name string;name = "tom"
	name := "tom"
	fmt.Println("name=",name)
}