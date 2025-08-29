package main

import (
	"fmt"
)

func main() {
	// ----------------- listas -----------------
	lista := []int{4, 9, 6, 7}
	fmt.Println("Lista: ", lista)

	for i := 0; i < len(lista); i++ {
		fmt.Println("Lista[i]: ", lista[i])
	}
	fmt.Println("Tamanho da lista: ", len(lista))

	lista = append(lista, 8)
	fmt.Println("Lista: ", lista)
	fmt.Println("Tamanho da lista: ", len(lista))

	listaDeString := []string{"golang", "c#", "java"}
	fmt.Println("ListaDeStrings: ", listaDeString)
	listaDeString = append(listaDeString, "ruby")
	fmt.Println("ListaDeStrings: ", listaDeString)

	// ----------------- make e for -----------------
	lista2 := make([]int, 1) // criando sem ser inicializado (precisa passar um tamanho inicial)
	lista2 = append(lista2, 3)
	lista2 = append(lista2, 3)
	fmt.Printf("%T\n", lista2)

	somaTotal := 0
	for i := 0; i < len(lista2); i++ {
		somaTotal += lista2[i]
	}
	fmt.Println("Média: ", (somaTotal / len(lista2)))

	// ----------------- sublistas e index -----------------
	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3}
	listaMenorQueCinco := make([]int, 0)

	for i := 0; i < len(listaToda); i++ {
		if listaToda[i] < 5 {
			listaMenorQueCinco = append(listaMenorQueCinco, listaToda[i])
		}
	}

	fmt.Println(listaMenorQueCinco)

	var segundaLista = listaToda[:3]
	var terceiraLista = listaToda[4:]
	fmt.Println(segundaLista)
	fmt.Println(terceiraLista)

	// ----------------- array vs slice -----------------
	/*
		no array é precisa informar um tamanho maximo que não pode ser aumentado, já no slice e dianmico e pode ser aumentado
		listaToda := [10]int{2, 10, 9, 4, 5, 8, 1, 3, 11, 22}	-> array
		listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3, 11, 22}		-> slice

		- array é muito mais perfomatico
		- slice é dinamico
	*/

	// ----------------- maps -----------------
	m := map[string]int{"sp": 100_000_000, "cg": 900_000} // modo de criar com valores
	fmt.Println(m)

	for chave, valor := range m {
		fmt.Println("Cidade:", chave, "\nHabitantes:", valor)
	}

	m2 := make(map[string]int) // modo de criar um map vazio
	m2["sp"] = 100_000_000
	m2["cg"] = 900_000
	m2["cg"] = 700_000
	m2["rj"] = 6_000_000

	valor, foiEncontrado := m2["rj"]
	if foiEncontrado {
		fmt.Println(valor)
	} else {
		fmt.Println("chave não existe")
	}

	/*
		-> OUTRA FORMA, USANDO A TRATATIVA JÁ NO IF

		if valor, foiEncontrado := m2["rj"]; foiEncontrado {
			fmt.Println(valor)
		} else {
			fmt.Println("chave não existe")
		}
	*/

	fmt.Println(m2)

	delete(m2, "cg")
	fmt.Println(m2)

}
