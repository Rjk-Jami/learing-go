package main

import (
	"fmt"
	"myPacakge/mathlib"
)

func sum(a int, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func printWecomeMessage() {
	fmt.Println("welcome to the application ")
}

func getUserName() string {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Printf("Enter your 1st number: ")
	fmt.Scanln(&num1)
	fmt.Printf("Enter your 2nd number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func displayResult(name string, sum int) {
	fmt.Println("Name:", name)
	fmt.Println("Sum:", sum)
}

func printGoodbyeMessage() {
	fmt.Println("Thank you! for the using our application")
}

func main() {
	// fmt.Println("Hello, Welcome to Go Programming ")
	// var car string = "BMW"
	// fmt.Println(car)

	// a := 10

	// b := "20"

	// a = 2.2

	// name := 30
	// fmt.Println(name)
	// fmt.Println(b)
	// fmt.Println(a)

	// sum(10, 20)

	// add := sum(10, 20)
	// fmt.Println("sum:", add)

	// a := 10
	// b := 20

	// p, q := getNumbers(a, b)

	// fmt.Println("sum:", p)
	// fmt.Println("mul:", q)

	// fmt.Println("welcome to the application ")

	// var name string
	// fmt.Printf("Enter your name: ")
	// fmt.Scanln(&name)

	// fmt.Println("Name:", name)

	// name := getUserName()
	num1, num2 := getTwoNumbers()

	// sum := sum(num1, num2)
	// displayResult(name, sum)
	// printGoodbyeMessage()

	sum := mathlib.Summation(num1, num2)
	fmt.Println("Sum:", sum)

}
