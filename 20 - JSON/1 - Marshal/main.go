package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Pitbull", 5}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON) //Bytes

	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) //Bytes para JSON

	c2 := map[string]string{
		"nome": "Cleito",
		"raca": "alemao",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON) //Bytes

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON)) //Bytes para JSON

}
