package model

import (
	"time"
)

type People struct {
	Nome       string
	Endereco   Adress
	Nascimento time.Time
	Idade      int
}

/*
func CalcYear(p People) int {
	yearNasc := p.Nascimento.Year()
	yearActual := time.Now().Year()
	return yearActual - yearNasc
}
*/

// exemplo de metodo
func (p *People) CalculaIdade() {
	yearNasc := p.Nascimento.Year()
	yearActual := time.Now().Year()
	p.Idade = yearActual - yearNasc
}
