package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan string)
	go print("hello", in)

	// for {
	// 	rcv, open := <-in
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(rcv)
	// }

	//The above algo replaced by
	for rcv := range in {
		fmt.Println(rcv)
	}
}

func print(msg string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- msg
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
