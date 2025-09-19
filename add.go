package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	println(add(2, 3))
	fmt.Println(add(10, 20))
