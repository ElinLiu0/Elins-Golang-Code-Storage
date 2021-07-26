package main

import "fmt"

func main() {
	//n换行
	fmt.Println("Hello\nWorld")
	//t空格
	fmt.Print("Hello\tWorld")
	fmt.Println("\nTom Says:\"Hi\"")
	// \r回车，从当前行的最前前面开始输出，覆盖掉以前的内容
	fmt.Println("My Name is\rElin")
	fmt.Println("姓名\t年龄\t籍贯\t住址\nJohn\t12\t河北\t北京")
	var num = 2 * 3
	fmt.Println("Result is %d",
		num * 2)
}
