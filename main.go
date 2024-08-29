package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to gobank!")
	storage, err := NewPostgreSQLStorage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", storage)
	server := NewAPIServer(":3000", storage)
	server.Run()
}
