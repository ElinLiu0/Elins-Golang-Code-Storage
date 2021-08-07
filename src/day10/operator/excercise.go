package main

import "fmt"

func main() {
	//练习1
	var days int = 97
	var week int
	var left_days int
	week = days / 7
	left_days = days % 7
	fmt.Printf("Total Working Weeks %d\n", week)
	fmt.Printf("Left Days %d\n", left_days)
	//练习2
	var F float64
	fmt.Scanf("Please Input The °F Temp%f\n", &F)
	var C float64
	C = 5.0 / 9 * (F - 100)
	fmt.Printf("After Changed into °C is %f\n", C)
}
