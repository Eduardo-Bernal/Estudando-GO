package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ---------------------------
	// ARRAYS
	// ---------------------------

	// Declarando um array de tamanho fixo 5 do tipo string.
	// Arrays em Go têm tamanho fixo e não podem ser alterados.
	var array1 [5]string

	// Atribui valor à posição 0 do array.
	array1[0] = "Posição 1"
	fmt.Println(array1) // Mostra o array completo (as posições vazias mostram "")

	// Declarando um array e inicializando todas as posições.
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Usando [...] o compilador descobre o tamanho do array automaticamente.
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// ---------------------------
	// SLICES
	// ---------------------------

	// Slice é como um array dinâmico (não tem tamanho fixo).
	// Aqui já criamos o slice com valores.
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// append adiciona um novo valor ao slice.
	// Slices podem crescer dinamicamente diferente de arrays.
	slice = append(slice, 18)
	fmt.Println(slice)

	// Criando um slice baseado em parte de um array:
	// Pega as posições 1 até 3 (posição final NÃO incluída).
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Alterar o array original altera o slice, porque slices apontam para o mesmo array interno.
	array2[1] = "Posição Alterada"
	fmt.Println(slice2) // Agora slice2 reflete essa mudança

	// ---------------------------
	// ARRAYS INTERNOS E MAKE
	// ---------------------------

	fmt.Println("----------")

	// Criando um slice usando make:
	// make(tipo, tamanho_inicial, capacidade)
	// tamanho = quantos elementos já existem
	// capacidade = quanto cabe antes de realocar
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3) // Slice com 10 posições preenchidas com 0

	// Inserindo valores com append
	slice3 = append(slice3, 5) // Agora tem 11 elementos → capacidade ainda cabe
	slice3 = append(slice3, 6) // Agora ultrapassa e o Go cria um novo array interno
	// Capacidade costuma dobrar quando precisa expandir

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // len = quantidade de elementos
	fmt.Println(cap(slice3)) // cap = capacidade interna do slice

	// Criando slice com make informando só tamanho (capacidade = tamanho)
	slice4 := make([]float32, 5)
	fmt.Println(slice4)

	// append aumenta o tamanho e provavelmente a capacidade
	slice4 = append(slice4, 10)

	fmt.Println(len(slice4)) // Agora é 6
	fmt.Println(cap(slice4)) // Capacidade provavelmente dobra para 10

}
