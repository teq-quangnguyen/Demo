package main

import "fmt"

func main() {
	go fmt.Println("======")
	go fmt.Println("hello")
	go fmt.Println("world")
	go fmt.Println("======")

	fmt.Println("todo")
}
