package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	//Formato Correto
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)

	//Formato Errado
	erro2 := checkmail.ValidateFormat("123")
	fmt.Println(erro2)
}

//Se a funcao comeca com letra minuscula ele apenas e reconhecida no pacote dela
