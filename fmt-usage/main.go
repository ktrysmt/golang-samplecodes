package main

import (
	"fmt"
)

func main() {
	verbs()
	widthAndPrecision()
}

func verbs() {
	fmt.Println("==verbs==")
	fmt.Printf("v:\t%v\n", "AAA")
	fmt.Printf("x:\t%x\n", 125)
	fmt.Printf("X:\t%X\n", 125)
	fmt.Printf("d:\t%d\n", 9)
	fmt.Printf("f:\t%f\n", 1.025)
	fmt.Printf("q:\t%q\n", " This is a pen. ")
	fmt.Printf("#v:\t%#v\n", " This is a pen. ")
	fmt.Printf("v:\t%v\n", " This is a pen. ")
	fmt.Printf("v:\t%v\n", 0x0022)
	fmt.Printf("q:\t%q\n", 00022)
	fmt.Printf("q,q:\t%q%q\n", 0x4b, 0x4f)
}
func widthAndPrecision() {
	fmt.Println("==widthAndPrecision==")
}
