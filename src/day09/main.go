package main

import (
	"fmt"

	testPtr "./functions"
)

func main() {
	testPtr.Test()
	fmt.Println(testPtr.Another)
}
