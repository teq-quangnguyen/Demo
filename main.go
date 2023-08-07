package main

import "fmt"

func main() {
	go fmt.Println("hello")
	go fmt.Println("world")
}
