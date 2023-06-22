package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do(a) Usuário(a) %s no banco de dados\n", u.nome)
}

func (u usuario) maioridade() bool {
	return u.idade >= 18
}

func (u *usuario) aniver() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	usuario1 := usuario{"Maria", 2}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Zayan", 1}
	fmt.Println(usuario2)
	usuario2.salvar()

	idade := usuario2.maioridade()
	fmt.Println(idade)
	idade2 := usuario1.maioridade()
	fmt.Println(idade2)

	usuario2.aniver()
	fmt.Println(usuario2)
	usuario1.aniver()
	fmt.Println(usuario1)
}
