package main

import (
	"fmt"
	"os"

	"github.com/jadercampos/go-concurrent-programming/goroutines"
)

func main() {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "1":
			goroutines.Before()
		case "2":
			goroutines.After()
		default:
			fmt.Println("Execute o programa seguido de um número de 1 a 9 para rodar os exercícios!")
		}
		fmt.Println(arg)
	}
}
