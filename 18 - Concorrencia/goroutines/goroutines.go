package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Paralelismo
	go escrever("Ola mundo") // GO ROUTINE Anula Concorrencia
	escrever("Programando em GO")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
