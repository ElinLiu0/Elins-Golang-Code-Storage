package main

import (
	"fmt"
	"strconv"
)

//演示Go中的字符串转成基本数据类型
func main(){
	var str string = "true"
	var b bool
	//说明：
	//1.strconv.ParseBool(str)函数会返回两个值（value bool,err error)
	//2.因为返回两个值，但是在Error我们并不关心，因此使用_忽略其返回值
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T and valus is '%v'\nm",b,b)
	//strconv.ParseInt(str,base,bitsize)
	//返回值：i int64,err error
	var str2 = "12345"
	var num1 int64
	num1 , _ = strconv.ParseInt(str2,10,0)
	fmt.Printf("num1 type %T and values is '%v'\n",num1,num1)
	var str3 string = "123.456"
	var num2 float64
	num2 , _ = strconv.ParseFloat(str3,64)
	fmt.Printf("num2 type %T and values is  '%v'\n",num2,num2)
	//注意：当Golang将字符串强制转换为整数值时会置零
	//强制转换为布尔值为false
	var str4 string = "hello"
	var num5 int64
	num5 , _ = strconv.ParseInt(str4,10,64)
	fmt.Printf("num5 type is %T and value is %v\n",num5,num5)
}
