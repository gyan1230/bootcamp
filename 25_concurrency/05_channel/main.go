package main

import "fmt"

func main() {
	ch := make(chan int)
	// send := make(chan<- int)
	// rcv := make(<-chan int)

	go f1(10, ch)
	n := <-ch
	fmt.Println(n)
}

func f1(n int, c chan int) {
	c <- n
}
