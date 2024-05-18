package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Interfaces:")

	retangulo := retangulo{
		largura: 3,
		altura:  2,
	}

	circulo := circulo{
		radius: 3,
	}

	Geometria(retangulo)
	Geometria(circulo)

}

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func Geometria(g geometria) {
	fmt.Println("Area:", g.area())
}
