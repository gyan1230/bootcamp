package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// go func() {
	// 	for {
	// 		fmt.Println(<-c) // this is a blocking call but for loop required
	// 	}

	// }()

	//above code replace by
	for v := range c { //this a blocking call
		fmt.Println(v)
	}
	//time.Sleep(1 * time.Second)	//no need to wait

}
