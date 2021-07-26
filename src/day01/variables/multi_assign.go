package main

import "fmt"
func main(){
	//该案例演示了GoLang语言中的多变量声明
	//法1
	//var n1,n2,n3 int
	//n1 = 1
	//n2 = 2
	//n3 = 3
	//var sum = n1 + n2 + n3
	//fmt.Println(sum)
	////法2
	//var nn1 , name , nn3 = 100 , "Elin" ,888
	//fmt.Println("nn1=",nn1,"name=",name,"nn3",nn3)
	//法3：适用类型推倒
	nn1 , name , nn3 := 100 , "Elin" ,888
	fmt.Println("nn1=",nn1,"name=",name,"nn3",nn3)
}