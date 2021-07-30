package main

import "fmt"

func main()  {
	var text string = "Hello Golang!"
	fmt.Println(text)
	//GoLang中的字符串是不可变的
	//var str = "hello"
	//此处无法为第0位元素更改内容str[0] = 'a'
	//输出源代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)
	str3 := `package main

import "fmt"

func main()  {
	var text string = "Hello Golang!"
	fmt.Println(text)
	//GoLang中的字符串是不可变的
	//var str = "hello"
	//此处无法为第0位元素更改内容str[0] = 'a'
	//输出源代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)`
	fmt.Println(str3)
	//字符串拼接方式
	var merge = "hello"+"\t"+"world"
	fmt.Println(merge)
	//当一个拼接的操作很长时，可以分行写，但是+留在被合并字符串旁边
}