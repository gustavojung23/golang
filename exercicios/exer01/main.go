package main

import (
	"exer01/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Compras Mercado")

	var nomeItem []string
	nomeItem = append(nomeItem, "Arroz")
	nomeItem = append(nomeItem, "Feijão")
	nomeItem = append(nomeItem, "Carne")
	nomeItem = append(nomeItem, "Ovo")

	compra, error := model.NewBuy("Mercado x", time.Now(), nomeItem)

	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println(compra)
	}
}
