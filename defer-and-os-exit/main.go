package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(test())
}

func test() int {
	defer deferTest()
	return 0
}

func deferTest() {
	fmt.Println("call as deferd")
}
