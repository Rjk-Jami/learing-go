package main

import (
	"fmt"
)

var a = 10

func main() {
	fmt.Println("Hello World")
	// age := 30
	// if age > 18 {
	// 	a := 47
	// 	fmt.Println(a)
	// }

	fmt.Println(a)

	// Anonymous function
	// Immediately invoked function expression (IIFE)

	// func(b int, c int) {
	// 	fmt.Println(b + c)
	// }(2, 3)

	// variadic function
	add := func(b int, c int) {
		fmt.Println(b + c)
	}
	add(2, a)

}

// inti function

func init() {
	fmt.Println("init")
	fmt.Println(a)
	// a = 20
}
