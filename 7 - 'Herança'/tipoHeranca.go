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
	matricula uint
}

func main() {
	fmt.Println("''Herança''")

	p1 := pessoa{"Jorge", "Antônio", 23, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "ADM", "Ulbra", 9578913}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.sobrenome)
	fmt.Println(e1.idade)
	fmt.Println(e1.altura)
	fmt.Println(e1.curso)
	fmt.Println(e1.faculdade)
	fmt.Println(e1.matricula)
}
