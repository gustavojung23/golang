package model

//herança

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
	Cor    string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	Portas         int
	Motor          float64
	ArCondicionado bool
}
