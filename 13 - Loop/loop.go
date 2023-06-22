package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loop")

	/*i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)
	fmt.Println("\n-----------------------\n")
	for j := 0; j < 10; j++ { //assim como no IF, variáveis declaradas dentro do for só funcionam no escopo do for em questão
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}*/

	nomes := [3]string{"Jeremy", "Joseph", "Karolyne"}

	for indice, nome := range nomes {
		indice++
		fmt.Println(indice, nome)
	}

	fmt.Println("\n-----------------------")

	for indice, letra := range "COMUNISMO" {
		indice++
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":  "Brendon",
		"idade": "24 Anos",
	}

	fmt.Println("\n-----------------------")

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não é possível usar o for com range em structs, ou seja, não se pode iterar as structs
}
