package main

import "fmt"

func funcao1() {

}
func funcao2() {

}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado sera retornado")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1()
	//Defer = ADIAR
	funcao2()

	fmt.Println(AlunoEstaAprovado(7, 8))
}
