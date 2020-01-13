package main

import "time"

import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "500 milisec"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "2 sec"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			{
				fmt.Println(msg1)
			}
		case msg2 := <-c2:
			{
				fmt.Println(msg2)
			}
		}
	}
}
