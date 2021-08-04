package main

import "fmt"

func main()  {
	//课堂练习1
	var num int = 10;
	var ptr *int = &num
	fmt.Printf("num的内存地址为：%v",ptr)
	//课堂练习2：通过ptr指针变更num内存下的值
	*ptr = 10
	fmt.Printf("\nnum修改后的值为：%v",num)
}
