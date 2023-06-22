package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	/*
		ARRAYS
	*/
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array3)

	/*
		SLICE
	*/
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 31)
	fmt.Println(slice)

	/*
		"INTERAÇÃO" ARRAY E SLICE
	*/
	slice2 := array2[1:4]
	fmt.Println(slice2)

	array2[1] = "Posição Modificada"
	fmt.Println(slice2)

	/*
		ARRAYS INTERNOS
	*/
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
