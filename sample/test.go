package main

import "fmt"
func main() {
	defer fmt.Println("World") // Delay print: "World"
	fmt.Println("Hello") // Print: "Hello"
	// output : Hello World
}
