package main

import (
	"fmt"
	"os"
)

func showText() {
	fmt.Println("Finalizando de manipular o arquivo.")
}

func main() {
	// =-=-=--=-=-=-=-=-=-== panic =-=-=--=-=-=-=-=-=-==
	/*
		Panic encerra o programa e retorna uma mensagem de erro
	*/

	_, err := os.Open("f:/settings.txt")
	if err != nil {
		panic(err)
	}

	// =-=-=--=-=-=-=-=-=-== defer =-=-=--=-=-=-=-=-=-==
	/*
		Defer vai ser a ultima coisa a ser execultada na função. neste caso depois de rodar toda a função o defer vai rodar o file.Close() para fechar o arquivo e depois vai ser execultado o outro defer que vai chamar a função showText
	*/

	file, err := os.Create("./settings.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
		
	defer showText()

	file.Write([]byte("teste"))

	// =-=-=--=-=-=-=-=-=-== recover =-=-=--=-=-=-=-=-=-==
	/*
		Recover não é uma boa pratica mas meio que vc tenta tratar o erro do panic
	*/

	ReadFile()
	fmt.Println("fim.")
}

func ReadFile() {
	defer func() { // tratando o erro do panic (função anonima)
		if r:= recover(); r != nil {
			fmt.Println("recuperado.")
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("FileNotExist")
	}
}