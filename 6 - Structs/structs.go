package main

import "fmt"

// Structs (algo parecido com classes em C#)
type usuario struct {
	nome  string
	idade int8
}

type endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("Arquivos structs")

	// Maneira tradicional
	var u usuario
	u.nome = "Eduardo"
	u.idade = 18
	fmt.Println(u)

	// Criando struct já preenchido
	enderecoExemplo := endereco{
		rua:    "Rua Tobias Barreto",
		numero: 0,
	}
	fmt.Println(enderecoExemplo)

	// Outra forma válida de criar struct (ordem dos campos!)
	usuario2 := usuario{"Eduardo", 18}
	fmt.Println(usuario2)

	// Quando quer definir apenas alguns campos
	usuario3 := usuario{idade: 18}
	fmt.Println(usuario3)
}
