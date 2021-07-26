package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//通过设置随机种子使得Go在每一次生成随机数是刷新内存
	rand.Seed(time.Now().Unix())
	fmt.Println("The Random Number is",rand.Intn(10));
	fmt.Println("The Random Number is",rand.Intn(10))
}
