package main

import "fmt"

func main() {
	/*
		Exercicio 01

		Criar um array com 2 posições de inteiros e armazenar em uma variavel a soma total da lista.
		A variavel deve ser imprimida no consoloe
	*/

	fmt.Println("Exercicio 01")

	lista := []int{2, 3}
	var soma int

	for i := range len(lista) {
		soma += lista[i]
	}

	fmt.Println("Soma:", soma)

	/*
		Exercicio 02

			Dado um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1" que vão de 1 a 10, efetuar a soma em duas variaveis, a primeira de numeros 1 a 5 e a segunda de 6 a 10.
			Imprimir os dois resultados
	*/

	fmt.Println("\nExercicio 02")

	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	var somaValoresMenores int
	var somaValoresMaiores int

	for i := range len(slice) {

		if slice[i] <= 5 {
			somaValoresMenores += slice[i]
		} else {
			somaValoresMaiores += slice[i]
		}
	}

	println("Soma dos valores de 1 a 5:", somaValoresMenores)
	println("Soma dos valores de 6 a 10:", somaValoresMaiores)
}
