package usingstatements

import (
	"fmt"
	"sync"
	"time"

	"github.com/jadercampos/go-concurrent-programming/entities"
)

func RunIfStatements() {
	fmt.Println("<<< RUN CHANNELS AND IF STATEMENTS >>>")
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func RunLoopingOver() {
	fmt.Println("<<< RUN LOOPING OVER CHANNELS >>>")
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func RunSelectStatements() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan entities.Book)
	dbCh := make(chan entities.Book)

	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- entities.Book) {
			if b, ok := entities.QueryCacheRWMutex(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- entities.Book) {
			if b, ok := entities.QueryDatabase(id); ok {
				m.Lock()
				entities.Cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan entities.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}
