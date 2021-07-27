package main

import "fmt"

var n1 = 200
var n2 = 100
var name = "Elin"

//上述方式可以被替换为：
var(
	n3 = 300
	n4 = 400
	name2 = "Jassie"
)
func main()  {
	fmt.Println("n1=",n1,"n2",n2,"name=",name)
	fmt.Println("n3=",n3,"n4=",n4,"name2=",name2)
}