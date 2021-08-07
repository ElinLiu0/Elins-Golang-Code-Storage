package main

import "fmt"

func main() {
	//重点讲解 /和%
	/*说明：如果运算的数都是整数，
	那么相除的结果自动去掉小数部分，
	并保留整数部分*/
	fmt.Println(10 / 4)
	/*之所以会输出值为2是因为10和4均为int
	而Go编译器为了保证数据结构一致因此输出值
	的类型也应当为int，所以输出值为2
	*/
	/*将值添加.0后输出正常*/
	fmt.Println(10.0 / 4.0)
	fmt.Println(10 % 3)
	/*++和--的使用*/
	var i int = 10
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
