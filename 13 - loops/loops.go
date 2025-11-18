package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("i =", i)
		// time.Sleep(time.Second) faz o programa "pausar"
		// por 1 segundo antes de continuar o loop.
		// Ou seja, a cada incremento de i o programa espera 1s.
		time.Sleep(time.Second)
	}

	fmt.Println("Valor final =", i)

	for j := 0; j < 64; j += 8 {
		j++
		fmt.Println("j =", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Eduardo", "Mayara", "Ganley"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Eduardo",
		"sobrenome": "Bernal",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//LOOP INFINITO
	for {
		fmt.Println("Executando i")
		time.Sleep(time.Second)
	}
}
