package waitgroups

import (
	"fmt"
	"sync"
	"time"

	"github.com/jadercampos/go-concurrent-programming/entities"
)

func Run() {
	fmt.Println("<<< RUN WAIT GROUPS >>>")
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := entities.QueryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := entities.QueryDatabase(id); ok {
				fmt.Println("from database")
				entities.Cache[id] = b
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}
