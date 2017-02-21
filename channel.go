package main

import "fmt"

func sum(s []int, c chan int, name string) {
	sum := 0
	for _, v := range s {
		fmt.Println(name, v)
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(s)
	c := make(chan int)
	go sum(s[:len(s)/2], c, "first")
	go sum(s[len(s)/2:], c, "second")
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
