package main

import (
	"fmt"
	"sync"
)

func main() {

	var ch = make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		ch <- 2
		ch <- 3
	}()
	wg.Wait()
	close(ch)
	for mm := range ch {
		fmt.Println(mm)
	}

}
