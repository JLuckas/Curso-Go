package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return (c.raio * c.raio) * math.Pi
}

func escreveArea(f forma) {
	fmt.Printf("A área da forma é %0.2fm²\n", f.area())
}

func main() {
	r := retangulo{15, 20}
	escreveArea(r)

	c := circulo{10}
	escreveArea(c)
}
