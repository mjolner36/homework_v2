package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1, chan2, chan3 := make(chan int), make(chan int), make(chan int)
	go func() {
		chan1 <- 1
		chan1 <- 3
		close(chan1)
	}()

	go func() {
		chan2 <- 2
		chan2 <- 4
		close(chan2)
	}()

	go func() {
		chan3 <- 0
		chan3 <- 5
		close(chan3)
	}()

	for ch := range merge(chan1, chan2, chan3) {
		fmt.Println(ch)
	}
}

func merge(channels ...chan int) chan int {
	resultChannel := make(chan int, 1)

	wg := &sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for c := range ch {
				resultChannel <- c
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	return resultChannel
}
