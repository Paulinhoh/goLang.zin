package main

import (
	"fmt"
)

func main() {
	// ------------------ Exemplo 01 ------------------
	var ehCarro bool = false
	var valorDoAutomovel = 1000.00

	if ehCarro {
		valorDoAutomovel += 55.50
	} else {
		valorDoAutomovel += 20.50
	}

	fmt.Println("Valor final do Automovel: ", valorDoAutomovel)
	fmt.Println()

	/*
		>  maior
		<  menor
		>= maio ou igual
		<= menor ou igula
		== igual
		!= diferente
	*/

	// ------------------ Exemplo 02 ------------------
	salario := 900.00
	tipo := "PJ"

	if salario < 1000.00 && tipo == "CLT" {
		salario -= (salario * 0.08)
	} else if salario < 1000.00 && tipo == "PJ" {
		salario -= (salario * 0.05)
	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.15)
	}
	fmt.Println("Salario final: ", salario)
	fmt.Println()

	/*
		Operadores lógicos
		&& and
		|| or
		!  not
	*/

	// ------------------ Exemplo 03 ------------------
	texto := "palavra"
	fmt.Println("Quantidade: ", len(texto))

	for i := 0; i <= len(texto); i++ {
		if string(texto[i]) == "r" {
			break // sai do laço
		}
		fmt.Println(string(texto[i]))
	}

	fmt.Println()

	// ------------------ Exemplo 04 ------------------
	texto2 := "palavra"
	tamanho := len(texto2)

	i := 0
	for i < tamanho {
		if string(texto2[i]) == "r" {
			i++
			continue // ele pula o resto da interação e volta para o laço
		}

		fmt.Println(string(texto2[i]))
		i++
	}

	fmt.Println()

	// ------------------ Exemplo 05 ------------------
	for numBase := 1; numBase <= 10; numBase++ {
		for multiplicado := 1; multiplicado <= 10; multiplicado++ {
			fmt.Println(numBase, " x ", multiplicado, " = ", (numBase * multiplicado))
		}
		fmt.Println()
	}

	fmt.Println()
}
