package main

import (
	"fmt"
	"time"
)

func func1 (s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go func1("world")
	func1("hello")
}