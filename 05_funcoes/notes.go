package main

import "fmt"

func main() {
	imprimeMensagem("mensagem x")
	imprimeMensagem("mensagem y")

	fmt.Println()

	resultado := soma(1, 2)
	fmt.Println("A soma é igual a:", resultado)

	fmt.Println()

	soma, subt, mult, div := operacoes(15, 5)
	fmt.Println(soma, subt, mult, div)

	fmt.Println()

	soma2, subt2, mult2, div2 := operacoes2(15, 5)
	fmt.Println(soma2, subt2, mult2, div2)
}

// funções
func imprimeMensagem(msg string) {
	msg += ", bom dia!"
	fmt.Println(msg)
}

func soma(num1 int, num2 int) int {
	resultado := num1 + num2
	return resultado
}

// função com retorno
func operacoes(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subt := num1 - num2
	mult := num1 * num2
	div := num1 / num2
	return soma, subt, mult, div
}

// retorno nomeado
func operacoes2(num1 int, num2 int) (soma int, subt int, mult int, div int) {
	soma = num1 + num2
	subt = num1 - num2
	mult = num1 * num2
	div = num1 / num2
	return
}

/*
	Escopo

		variaveis dentro de funções são de escopo local ou seja só funcionam e só podem ser acessadas dentro da propria função

		Variaveis que são criadas a nivel de arquivo ou seja fora das funções são variaveis globais e podem ser acessadas de qualquer lugar do arquivo

*/

// função init -> essa função init ela é chamada automaticamente quando o programa é iniciado
func init() {
	fmt.Println("Função Init 1")
	fmt.Println()
}

func init() {
	fmt.Println("Função Init 2")
	fmt.Println()
}
