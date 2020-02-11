package main

import (
	"fmt"
	"time"
)

// Synchronous code
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Concurrency with Go Routines")

	go compute(5)
	go compute(5)

	fmt.Scanln()
}
