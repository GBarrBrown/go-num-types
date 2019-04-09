package main

import "fmt"

func main() {
	printBools()
	printString()
}

func printBools() {
	fmt.Println("-- Boolean -- ")
	var isTrue bool = true
	var isFalse bool = false
	fmt.Println(isTrue)
	fmt.Println(isFalse)
}

func printString() {
	fmt.Println("-- String -- ")
	var stringOne string = "I'm a string type"
	fmt.Println(stringOne)
}