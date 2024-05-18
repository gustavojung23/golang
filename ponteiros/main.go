package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 10

	imprimirValores(y, &x)
	fmt.Println(*y, x)
	fmt.Println(y, &x) // & = mostra o valor da memória da variável
}

func imprimirValores(y *int, x *int) {
	*x = 20
}
