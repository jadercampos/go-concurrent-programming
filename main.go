package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jadercampos/go-concurrent-programming/async"
	bufferedchannels "github.com/jadercampos/go-concurrent-programming/buffered_channels"
	channeltypes "github.com/jadercampos/go-concurrent-programming/channel_types"
	closingchannels "github.com/jadercampos/go-concurrent-programming/closing_channels"
	goroutines "github.com/jadercampos/go-concurrent-programming/go_routines"
	"github.com/jadercampos/go-concurrent-programming/mutexes"
	rwmutexes "github.com/jadercampos/go-concurrent-programming/rw_mutexes"
	unbufferedchannels "github.com/jadercampos/go-concurrent-programming/unbuffered_channels"
	usingstatements "github.com/jadercampos/go-concurrent-programming/using_statements"
	waitgroups "github.com/jadercampos/go-concurrent-programming/wait_groups"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Execute o programa seguido de um número de 1 a 13 para rodar os exercícios!")
	} else {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "1":
				goroutines.Before()
			case "2":
				goroutines.After()
			case "3":
				waitgroups.Run()
			case "4":
				mutexes.Run()
			case "5":
				rwmutexes.Run()
			case "6":
				unbufferedchannels.Run()
			case "7":
				bufferedchannels.Run()
			case "8":
				channeltypes.Run()
			case "9":
				closingchannels.Run()
			case "10":
				usingstatements.RunIfStatements()
			case "11":
				usingstatements.RunLoopingOver()
			case "12":
				usingstatements.RunSelectStatements()
			case "13":
				AsyncExample()
			default:
				fmt.Println("Exercício informado inválido, apenas opções de 1 a 13")
			}
		}
	}

}

func AsyncExample() {
	fmt.Println("<<< RUN ASYNC EXAMPLE >>>")
	fmt.Println("Let's start ...")
	future := async.Exec(func() interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running ...")
	val := future.Await()
	fmt.Println("Done is done: ", val)
}

func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 69
}
