package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// ----------------- hello, world! -----------------
	fmt.Println("hello, world!")

	// ----------------- variaveis -----------------
	var peso int
	peso = 10
	var mensagem string = "10 de peso"
	mensagem2 := "10.5" // nesta forma não é necessario a palavra var nem a typagem da variavel
	fmt.Println(peso)
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

	// ----------------- tipos de dados -----------------
	/*
		No go a typagem é estatica, ou seja uma vez definida o tipo da variavel não pode ser mudado
	*/

	var texto string
	var numero int
	var metro float32
	var ehfeminino bool

	fmt.Println(texto)
	fmt.Println(numero)
	fmt.Println(metro)
	fmt.Println(ehfeminino)

	/*
		string
		int
		float
		bool
		...
	*/

	// ----------------- operadores -----------------
	num1 := 10.0
	num2 := 20.0
	result := num1 / num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

	text1 := "texto 1"
	text2 := "texto 2"
	result2 := text1 + text2
	fmt.Println(result2)
	fmt.Println(reflect.TypeOf(result2))

	/*
		+  soma
		-  subtração
		*  multiplicação
		/  divisão
		%  resto
	*/

	// ----------------- Constantes -----------------
	const taxa = 10

	// ----------------- Tamanho dos tipos de dados -----------------
	var numero2 uint8
	numero2 = 1
	fmt.Println(numero2)

	/*
		int8		8-bit signed integer
		int16		16-bit signed integer
		int32		32-bit signed integer
		int64		64-bit signed integer
		uint8		8-bit unsigned integer
		uint16		16-bit unsigned integer
		uint32		32-bit unsigned integer
		uint64		64-bit unsigned integer
		int			8oth in and uint contain same size, either 32 and 64 bit
		uint		8oth in and uint contain same size, either 32 and 64 bit

		float32		32-bit IEEE 754 floating-point number
		float64		64-bit IEEE 754 floating-point number
	*/

	// ----------------- Conversão -----------------
	var numero3 int8 = 127
	var numeroInt int
	numeroInt = int(numero3)
	fmt.Println(reflect.TypeOf(numeroInt))
	var numeroInt8 int8 = int8(numeroInt)
	fmt.Println(reflect.TypeOf(numeroInt8))
	var numeroFloat32 float32 = float32(numeroInt8)
	fmt.Println(reflect.TypeOf(numeroFloat32))

	// strconv -> pacote para converter strings
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(i)
	fmt.Printf("%T \n", i)

	text4 := "24.15"
	f, _ := strconv.ParseFloat(text4, 64)
	fmt.Println(f)
	fmt.Printf("%T \n", f)
}
