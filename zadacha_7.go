package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3) //исправил сделал канал буферезированым
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	close(ch)
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}

// на 9 строчке нужно сделал буферезированный канал
//ch := make(chan int) //не буферезированный в него будет записано только одно значение и горутина main блокируется wg.Wait()
