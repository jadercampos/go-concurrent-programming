package mutexes

import (
	"fmt"
	"sync"
	"time"

	"github.com/jadercampos/go-concurrent-programming/entities"
)

func Run() {
	fmt.Println("<<< RUN MUTEXES >>>")
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok := entities.QueryCacheMutex(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			if b, ok := entities.QueryDatabase(id); ok {
				fmt.Println("from database")
				m.Lock()
				entities.Cache[id] = b
				m.Unlock()
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}
