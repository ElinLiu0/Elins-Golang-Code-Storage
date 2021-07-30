package main

import "fmt"

func main()  {
	var a int
	var b float32
	var c float64
	var name string
	var dead bool
	fmt.Printf("a=%d\n",a)
	fmt.Printf("b=%f\n",b)
	fmt.Printf("c=%f\n",c)
	fmt.Printf("name=%s\n",name)
	fmt.Printf("dead=%v\n",dead)
	//在格式化输出时，%v表示输出变量的原始值进行输出
}
