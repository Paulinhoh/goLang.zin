package main

import "github.com/fatih/color"

func main() {
	// =-=-=-=-=-=-==--=-=-=-=-= pacotes =-=-=-=-=-=-==--=-=-=-=-=

	/*
		exemplo de instalação de um pacote externo

		1. inicia o go mod
		-> go mod init <nome do módulo>

		2. installar o pacote
		-> go get github.com/fatih/color (exemplo)

		3. baixar as dependecias
		-> go get
		-> (ou) go mod tidy
	*/

	color.Magenta("Mensagem em magenta")
}