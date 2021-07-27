package main

import "fmt"

//变量使用的注意事项
func main()  {
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=",i)
	//此时的i无法被为浮点数：因为声明时已被指定类型
	//i = 1.2
}
