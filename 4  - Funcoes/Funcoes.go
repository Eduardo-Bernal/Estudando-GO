package main

import (
	"fmt"
)

//Parametros passados dentro do () e depois tipo

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// Funcoes tambem sao tipos
// Podendo igualar uma variavel a uma funcao

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da Funcao 1")

	//Tipo 1 - com as 2 variaveis
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Tipo 2 - com as 1 variaveis
	//resultadoSoma, _ := calculosMatematicos(10, 15)
	//fmt.Println(resultadoSoma)

}
