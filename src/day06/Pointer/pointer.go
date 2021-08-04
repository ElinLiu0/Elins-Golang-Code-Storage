package main

import "fmt"

//演示GoLang中的指针类型		
func main()  {
	//基本数据类型在内存的布局
	var i int = 10
	// 如何取出i的内存地址
	fmt.Println("i的地址=",&i)
	//下面的var ptr *int = &i
	//1.ptr 是一个指针变量
	//2.ptr的类型是*int
	//3.ptr本身的值为&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr的地址=%v\n",&ptr)
	fmt.Printf("ptr 指向的值=%v\n",*ptr)
}