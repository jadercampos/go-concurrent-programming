package channeltypes

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println("<<< RUN CHANNEL TYPES>>>")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
