package main

import "fmt"

// int8, int16, int32, int64
// Apenas INT = Arquitetura da maquina por exemplo 64 bits
func main() {
	var numero int = -1000000000
	fmt.Println(numero)

	//uint = unsigned int - APENAS NUMEROS POSITIVOS || = 0
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias - apelido para se referir
	//Usados quando tem caracteres sem ser numeros
	//INT 32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	//INT 8 = Byte
	var numero4 rune = 1234
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456789
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Fim Numeros Reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Nao temos CHAR no GOLANG porem retorna valor
	//Da tabela ASC
	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	//Todo dado no GO comeca com valor inicial = 0
	var texto string
	fmt.Println(texto)

	//Boolean = True e False || 0 ou 1

	var Booleano1 bool
	fmt.Println(Booleano1)

	//Tipo Error no GO
	//Retorna nill
	var erro error
	fmt.Println(erro)

}
