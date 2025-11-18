package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	u1 := usuario{"Edu", 18}
	fmt.Println(u1)
	u1.salvar()
	u1.maiorDeIdade()
	MaiorDeIdade := u1.maiorDeIdade()
	fmt.Println(MaiorDeIdade)
	u1.fazerAniversario()
	fmt.Println("Parabens e seu aniversario agora voce tem", u1.idade)

	fmt.Println("-------------------------")

	u2 := usuario{"May", 17}
	fmt.Println(u2)
	u2.salvar()
	MaiorDeIdade = u2.maiorDeIdade()
	fmt.Println(MaiorDeIdade)
	u2.fazerAniversario()
	fmt.Println("Parabens e seu aniversario agora voce tem", u2.idade)
	MaiorDeIdade = u2.maiorDeIdade() //AGORA FICA DE MAIOR
	fmt.Println(MaiorDeIdade)

}
