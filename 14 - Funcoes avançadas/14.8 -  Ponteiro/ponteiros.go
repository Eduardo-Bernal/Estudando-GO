package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComponteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	NovoNumero := 8453
	fmt.Println(NovoNumero)
	inverterSinalComponteiro(&NovoNumero)
	fmt.Println(NovoNumero)

}
