package main

import "fmt"

//演示Golang中表示符的使用
func main() {
	//Golang中严格区分大小写
	var num int = 0
	var Num int = 100
	fmt.Printf("num = %d\n", num)
	fmt.Printf("Num = %d\n", Num)
	//标识符不能包含空格
	//var ab c int = 20,10
	// _是空标识符，用于占用
	// var _ int = 400
	// fmt.Printf(_)
}
