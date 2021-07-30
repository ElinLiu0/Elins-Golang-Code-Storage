package main

import "fmt"

func main()  {
	var c6 byte = 'a'
	var user byte
	for i :=1 ; i<=10;i++{
		fmt.Println("请输入一个值")
		fmt.Scan(&user)
		if user == c6 {
			fmt.Printf("猜对了")
			break
		}
		if i == 10{
			fmt.Println("次数用完了")
			break
		}
	}
}