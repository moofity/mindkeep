package main

import (
	"fmt"
	"time"
)

func main() {
	go printHello()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		printHello()
	}
}

func printHello() {
	fmt.Println("Hello, World!")
}
