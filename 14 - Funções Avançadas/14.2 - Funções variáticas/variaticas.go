package main

import "fmt"

func soma(numeros ...int) int { //Só pode ter apenas UM parametro variático
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) { //O parametro variático tem que ser o último parametro de uma função
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(totalSoma)

	escrever("Olá mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
}
