package main

import (
	"fmt"
	"enderecos/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Rua João Dutra Filho")
	fmt.Println(tipoEndereco)
}