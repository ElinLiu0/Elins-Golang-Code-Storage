package main

import "fmt"

func main() {
	var i int = 8
	var a int
	/*错误方法：a = i++
	说明++需要独立使用
	*/
	i++
	a = i
	fmt.Println(a)
}
