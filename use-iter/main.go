package main

import (
	"fmt"
)

func main() {
	fmt.Println("---")
	fmt.Println(num(9))

	for i := range num(9) {
		fmt.Println(i)
	}

	for i, v := range make([]int, 9) {
		fmt.Println(i, v)
	}
}

func num(n int) []int {
	return make([]int, n)
}
