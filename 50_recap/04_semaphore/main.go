package main

import "fmt"

func main() {
	ch := make(chan int)
	signal := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		signal <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		signal <- true

	}()

	go func() {
		<-signal
		<-signal
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}

}
