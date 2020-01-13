package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		print("hello")
		wg.Done()
	}()
	wg.Wait()
}

func print(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println("msg :", msg, "-", i)
		time.Sleep(time.Millisecond * 500)
	}
}
