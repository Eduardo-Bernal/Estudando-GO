package main

import (
	"fmt"
	"time"
)

// é um canal de comunicação o canal  enviar como receber dados está.
// Então a gente usa esse envio e recebimento de dados para sincronizar as nossas goroutines.
func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)

	fmt.Println("Depois da funcao")
	// for {
	// 	msg, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	//! OU
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
