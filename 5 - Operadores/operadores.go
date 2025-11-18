package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 10 + 5
	subtracao := 10 - 5
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int32 = 10
	var numero2 int32 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	//fim dos aritmeticos

	// OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//Fim dos operadores de atribuição

	// OPERADORES RELACIONAIS
	fmt.Println("RELACIONAIS:")
	fmt.Println(10 > 5)   // maior que
	fmt.Println(10 < 5)   // menor que
	fmt.Println(10 >= 5)  // maior ou igual
	fmt.Println(10 <= 5)  // menor ou igual
	fmt.Println(10 == 10) // igual
	fmt.Println(10 != 5)  // diferente
	//Fim dos operadores relacionais

	// OPERADORES LÓGICOS
	fmt.Println("LÓGICOS:")
	fmt.Println(true && false) // AND
	fmt.Println(true || false) // OR
	fmt.Println(!true)         // NOT
	//Fim dos operadores lógicos

	// OPERADORES UNÁRIOS
	fmt.Println("UNÁRIOS:")
	num := 10
	num++ // incremento
	fmt.Println(num)

	num-- // decremento
	fmt.Println(num)

	num += 5 // soma e atribui
	fmt.Println(num)

	num -= 3 // subtrai e atribui
	fmt.Println(num)
	//Fim dos operadores unários
}
