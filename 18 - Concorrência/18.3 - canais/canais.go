package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois função escrever!")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
