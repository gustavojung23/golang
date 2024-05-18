package main

import (
	"fmt"
	model "struct/model"
	"time"
)

func main() {
	fmt.Println("Inicio")

	endereco := model.Adress{
		Rua:    "Rua x",
		Numero: 10,
		Cidade: "Corbélia",
	}
	pessoa := model.People{
		Nome:       "Gustavo",
		Endereco:   endereco,
		Nascimento: time.Date(1995, 8, 23, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(endereco)
	//fmt.Println(endereco.Rua)
	//fmt.Println(endereco.Numero)
	//fmt.Println(endereco.Cidade)

	//endereco.Cidade = "Cascavel"
	//fmt.Println(endereco.Cidade)

	fmt.Println(pessoa)

	//year := model.CalcYear(pessoa)
	//year := pessoa.CalculaIdade()
	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)

	autoMovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "XPT-3452",
		Modelo: "CB 600",
		Cor:    "Preta",
	}

	moto := model.Moto{
		Automovel:   autoMovelMoto,
		Cilindradas: 650,
	}

	fmt.Println(moto)
}
