package main

import (
	"errors"
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

	// Errors
	Errors(errors.New("a error"))
	p := ProblemaNetwork{
		Rede:     true,
		Hardware: false,
	}
	Errors(p)

	//interface vazia
	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, 7.5)
	lista = append(lista, true)
	lista = append(lista, "teste")

	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		}
		fmt.Println(valor)
	}

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

// struct error
type ProblemaNetwork struct {
	Rede     bool
	Hardware bool
}

func (p ProblemaNetwork) Error() string {
	if p.Rede {
		return "Problema de Rede"
	} else if p.Hardware {
		return "Problema de Hardware"
	} else {
		return "Outro Problema"
	}
}

// func error
func Errors(err error) {
	fmt.Println(err.Error())
}
