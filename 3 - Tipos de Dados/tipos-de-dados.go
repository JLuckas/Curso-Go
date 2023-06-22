package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
	*
	*
	*		NÚMEROS INTEIROS
	*
	*
	 */

	fmt.Println("\n<========== Números Inteiros ==========>\n")
	// Tipos de Inteiros(Os sufixos de numero são referentes aos bits suportados): int8, int16, int32, int64
	numero := -1000000000000000000
	fmt.Println(numero)

	// uint = unsigned int
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//alias uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	/*
	*
	*
	*		NÚMEROS REAIS
	*
	*
	 */

	fmt.Println("\n<========== Números Reais ==========>\n")
	var numeroReal1 float32 = 123.456
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456789101112.13
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.678
	fmt.Println(numeroReal3)

	/*
	*
	*
	*		STRINGS
	*
	*
	 */

	fmt.Println("\n<========== Strings ==========>\n")
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	/*
	*
	*
	*		VALOR ZERO
	*
	*
	 */

	fmt.Println("\n<========== Valores Zero ==========>\n")
	var texto string
	fmt.Println(texto)

	var texto2 int
	fmt.Println(texto2)

	var texto3 error
	fmt.Println(texto3)

	/*
	*
	*
	*		BOOLEAN
	*
	*
	 */

	fmt.Println("\n<========== Boolean ==========>\n")
	var bool1 bool = true
	fmt.Println(bool1)

	var bool2 bool = false
	fmt.Println(bool2)

	/*
	*
	*
	*		ERROR
	*
	*
	 */
	fmt.Println("\n<========== Error ==========>\n")
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
