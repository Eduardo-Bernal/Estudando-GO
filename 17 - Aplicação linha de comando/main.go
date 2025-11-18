package main

import (
	"fmt"
	"linhadecomando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	Aplicacao := app.Gerar()
	if erro := Aplicacao.Run(os.Args); erro != nil {
		log.Fatal((erro))
	}

}
