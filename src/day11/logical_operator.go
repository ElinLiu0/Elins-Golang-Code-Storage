package main

import "fmt"

//声明一个测试函数
func test() bool {
	fmt.Println("test..")
	return true
}
func main() {
	/*短路与
	因为i==9为false，因此后面放弃判断
	*/
	var i int = 10
	if i == 9 && test() {
		fmt.Println("ok")
	}
	/*短路或
	因为i>9为true，因此后面放弃判断
	*/
	if i > 9 || test() {
		fmt.Println("either ok")
	}
}
