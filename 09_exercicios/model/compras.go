package model

import (
	"errors"
	"time"
)

type Compras struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) (*Compras, error) {
	if mercado == "" {
		return nil, errors.New("mercado é obrigatório")
	}
	if len(nomeDosItens) == 0 {
		return nil, errors.New("itens são obrigatórios")
	}

	var itens []ItemDaCompra
	for _, valor := range nomeDosItens {
		itens = append(itens, ItemDaCompra{Nome: valor})
	}

	compras := &Compras{
		Mercado: mercado,
		Data:    data,
		Itens:   itens,
	}

	return compras, nil
}
