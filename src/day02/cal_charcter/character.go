package main

import "fmt"

func main()  {
	//加法计算
	var(
		i = 1
		j = 2
		r = i + j
	)
	fmt.Println("r=",r)
	//字符串
	var(
		str1 = "hello"
		str2 = "world"
		result = str1+str2
	)
	fmt.Println("merged string is:",result)
}