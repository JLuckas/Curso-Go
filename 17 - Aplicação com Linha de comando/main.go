package main

import (
	"fmt"
	"linha-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Aplicação com Linha de comando")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
