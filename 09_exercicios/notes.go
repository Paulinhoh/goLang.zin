package main

import (
	"fmt"
	"golangestudos/09_exercicios/model"
	"time"
)

func main() {
	/*
		Exercicios
		- Criar um modelo que ira receber os itens para uma compra do mês, nesse modelo teremos a data que a compra irá acontecer, mercado e os itens para comprar
		- Dado o exercicio anterior, mover o modelo anterior criado para o pacote chamado "Model"
		- Dado o exercicio anterior, criar uma função no pacote "Model" que inicializa a struct e retorna como ponteiro
	*/

	var nomeDosItens []string
	// nomeDosItens = append(nomeDosItens, "Arroz")
	// nomeDosItens = append(nomeDosItens, "Feijão")
	// nomeDosItens = append(nomeDosItens, "Carne")
	// nomeDosItens = append(nomeDosItens, "Shampoo")

	compras, err := model.NewCompra("Mercas", time.Now(), nomeDosItens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compras)
	}
}
