package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var WaitGroup sync.WaitGroup

	WaitGroup.Add(2)

	go func() {
		go escrever("Ola mundo")
		WaitGroup.Done()
	}()
	go func() {
		escrever("Programando em GO")
		WaitGroup.Done()
	}()

	WaitGroup.Wait() // Executa 2 goroutines ate chegar em zero

}

func escrever(texto string) {
	for i := 0; i < 2; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
