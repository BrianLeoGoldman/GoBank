package main

import "fmt"

func main() {
	fmt.Println("Welcome to gobank!")
	server := NewAPIServer(":3000")
	server.Run()
}
