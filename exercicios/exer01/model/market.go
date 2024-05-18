package model

import (
	"errors"
	"time"
)

type Itens struct {
	NomeItem string
}

type Market struct {
	Mercado string
	Data    time.Time
	Item    []Itens
}

func NewBuy(mercado string, data time.Time, nomeItem []string) (*Market, error) {

	if mercado == "" {
		return nil, errors.New("Mercado é obrgatório")
	}

	if len(nomeItem) == 0 {
		return nil, errors.New("Itens é obrgatório")
	}

	var itens []Itens

	for _, produto := range nomeItem {
		itens = append(itens, Itens{NomeItem: produto})
	}

	return &Market{
		Mercado: mercado,
		Data:    time.Now(),
		Item:    itens,
	}, nil
}
