package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//Uma variável declarada dentro de um if funciona apenas dentro do escopo do if em questão
	if numero1 := numero; numero1 > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero1 < -5 {
		fmt.Println("O número é menor que 0 e menor que -5")
	} else {
		fmt.Println("O número é menor que 0")
	}

}
