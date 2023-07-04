package main

import (
	"fmt"
	"enconding/json"
	"log"
)

type cachorro struct {
	nome string `json:"nome"`
	raca string `json:"raca"`
	idade uint `json:"idade"`
}

func main() {
	c := cachorro{"Pernalong", "Vira-Lata", 9}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)
}