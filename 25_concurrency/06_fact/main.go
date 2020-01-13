package main

import "fmt"

func fact(n uint64, c chan uint64) {
	var f, i uint64
	f = 1
	for i = 2; i < n; i++ {
		f = f * i
	}
	c <- f
}

func main() {
	ch := make(chan uint64)
	defer close(ch)

	for i := 1; i < 15; i++ {
		go fact(uint64(i), ch)
		fmt.Println(<-ch)
	}
}
