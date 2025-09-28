package main

import (
	"fmt"
	"golangestudos/08_struct/model"
	"time"
)

/*
type endereco struct {
	rua    string
	numero int
	cidade string
}
*/

func main() {
	// -=-=-=-=-=-=-=-=-=-=-=-=-= go mod -=-=-=-=-=-=-=-=-=-=-=-=-=
	/*
		comando para criar o arquivo go.mod: go mod init <nome do módulo>
		comando para baixar as dependencias: go mod tidy
	*/

	// -=-=-=-=-=-=-=-=-=-=-=-=-= struct -=-=-=-=-=-=-=-=-=-=-=-=-=
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Panificador José Silva",
		Numero: 46,
		Cidade: "São Cristovão",
	}
	fmt.Println(endereco)

	endereco.Numero = 22
	fmt.Println(endereco.Numero)

	// -=-=-=-=-=-=-=-=-=-=-=-=-= encapsulamento -=-=-=-=-=-=-=-=-=-=-=-=-=
	/*
		- Se começa com letra MAISCULA a struct é acessivel em outros pacotes
		- Se começa com letra minuscula é privado ou seja acessivel somente no prórpio pacote
	*/

	// -=-=-=-=-=-=-=-=-=-=-=-=-= composição -=-=-=-=-=-=-=-=-=-=-=-=-=
	pessoa := model.Pessoa{
		Nome:             "Paulo Henrique Reis",
		Endereco:         endereco,
		DataDeNascimento: time.Date(2001, 07, 31, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(pessoa)

	// -=-=-=-=-=-=-=-=-=-=-=-=-= métodos -=-=-=-=-=-=-=-=-=-=-=-=-=
	idade := model.CalculaIdade(pessoa)
	pessoa.CalculaIdade()
	fmt.Println(idade)
	fmt.Println(pessoa.Idade)

	// -=-=-=-=-=-=-=-=-=-=-=-=-= alterando dados por métodos -=-=-=-=-=-=-=-=-=-=-=-=-=
	/*
		- Para alterar algum atributo em algum método deve ser usado ponteiros se não ele ira criar outro espaço na memoria para a realização do método e não ira ser salvo na mémoria correta
	*/
	// -=-=-=-=-=-=-=-=-=-=-=-=-= herança -=-=-=-=-=-=-=-=-=-=-=-=-=

	automovelMoto := model.Automovel {
		Ano: 2022,
		Placa: "XPTO",
		Modelo: "GC",
	}

	moto := model.Moto {
		Automovel: automovelMoto,
		Cilindradas: 125,
	}
	fmt.Println(moto)
	fmt.Println(moto.Modelo)
}
