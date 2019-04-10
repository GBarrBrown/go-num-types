package main

import "fmt"

func main() {
	printBools()
	printString()
	printInts()
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

func printInts() {
	fmt.Println("-- Integer --")
	var isInt int = 6
	fmt.Println("int 6: ",isInt)
}