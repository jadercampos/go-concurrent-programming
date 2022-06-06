package main

import (
	"fmt"
	"os"

	"github.com/jadercampos/go-concurrent-programming/goroutines"
	"github.com/jadercampos/go-concurrent-programming/mutexes"
	"github.com/jadercampos/go-concurrent-programming/rwmutexes"
	"github.com/jadercampos/go-concurrent-programming/waitgroups"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Execute o programa seguido de um número de 1 a 9 para rodar os exercícios!")
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
			default:
				fmt.Println("Exercício informado inválido, apenas opções de 1 a 9")
			}
			fmt.Println(arg)
		}
	}

}
