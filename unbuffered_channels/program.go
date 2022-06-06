package unbuferredchannels

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println("<<< RUN UNBUFFERED CHANNELS >>>")
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	wg.Add(2)
	go func(ch chan int, ch2 chan string, ch3 chan bool, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch2)
		fmt.Println(<-ch3)
		wg.Done()
	}(ch, ch2, ch3, wg)
	go func(ch chan int, ch2 chan string, ch3 chan bool, wg *sync.WaitGroup) {
		ch <- 4915
		ch2 <- "Oi sumido S2 !"
		ch3 <- true
		wg.Done()
	}(ch, ch2, ch3, wg)

	wg.Wait()
}
