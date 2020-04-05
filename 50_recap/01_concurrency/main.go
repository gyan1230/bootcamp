package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
	fmt.Println("Count:", count)
}

func increment(rcv string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		mutex.Lock()
		x := count
		x++
		count = x
		fmt.Println(rcv, i, count)
		mutex.Unlock()
	}
	wg.Done()
}
