package main

import "fmt"

func recuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado!")
	}
}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecução()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!!")

}

func main() {
	fmt.Println(AlunoEstaAprovado(6, 6))
	fmt.Println("Pos execução!")
}
