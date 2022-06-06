package goroutines

import (
	"fmt"
	"time"

	"github.com/jadercampos/go-concurrent-programming/entities"
)

func Before() {
	fmt.Println("<<< BEFORE >>>")
	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		if b, ok := entities.QueryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := entities.QueryDatabase(id); ok {
			fmt.Println("from database")
			entities.Cache[id] = b
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
}
func After() {
	fmt.Println("<<< AFTER >>>")
	for i := 0; i < 10; i++ {
		id := entities.Rnd.Intn(10) + 1
		go func(id int) {
			if b, ok := entities.QueryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := entities.QueryDatabase(id); ok {
				entities.Cache[id] = b
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(id)
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
}
