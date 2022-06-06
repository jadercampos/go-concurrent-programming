package rwmutexes

import (
	"fmt"
	"sync"
	"time"

	"github.com/jadercampos/go-concurrent-programming/entities"
)

func Run() {
	fmt.Println("<<< RUN READ/WRITE MUTEXES >>>")
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		wg.Add(2)
		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := entities.QueryCacheRWMutex(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(wg, m)
		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := entities.QueryDatabase(id); ok {
				fmt.Println("from database")
				m.Lock()
				entities.Cache[id] = b
				m.Unlock()
				fmt.Println(b)
			}
			wg.Done()
		}(wg, m)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}
