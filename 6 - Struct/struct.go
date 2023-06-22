package main

import "fmt"

type usuario struct {
	name     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("STRUCTS")

	var u usuario
	u.name = "Joelison"
	u.idade = 16
	fmt.Println(u)

	enderedoExemplo := endereco{"Avenida Antônio Gomes Corrêa", 1500}
	u2 := usuario{"João", 22, enderedoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 36}
	fmt.Println(u3)
	u3 = usuario{name: "César"}
	fmt.Println(u3)

}
