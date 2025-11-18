package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("heranca")

	p1 := pessoa{"Eduardo", "Bernal", 18, 165}
	fmt.Println(p1)

	e1 := estudante{p1, "DEV", "SENAI"}
	fmt.Println(e1)
}
