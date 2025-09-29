package main

import "fmt"


func reverse(slice []int) []int {

	newInts := make([]int, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}

func reverse2(slice []string) []string {

	newStrings := make([]string, len(slice))

	newStringsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newStrings[newStringsLen] = slice[i]
		newStringsLen--
	}

	return newStrings
}

func main() {
	slice := []int{5, 1, 2, 3}
	slice2 := []string{"a", "e", "f", "b"}
	newInts := reverse(slice)
	newStrings := reverse2(slice2)

	fmt.Println(newInts)
	fmt.Println(newStrings)
	fmt.Println()

	// =-=-==-==-=-=-=-=-=-=-=-= Generics =-=-==-==-=-=-=-=-=-=-=-=
	/*
	- Interfaces resolvem problemas de comportamento
	- Generics resovem problemas de dados
	*/
	
	newIntsGeneric := reverseGeneric(slice)
	newStringsGeneric := reverseGeneric(slice2)
	newStringsGeneric2 := reverseGeneric2(slice2)
	
	fmt.Println(newIntsGeneric)
	fmt.Println(newStringsGeneric)
	fmt.Println(newStringsGeneric2)

	// =-=-==-==-=-=-=-=-=-=-=-= Constraints =-=-==-==-=-=-=-=-=-=-=-=
}

func reverseGeneric[T int | string](slice []T) []T {

	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}

type constraintCustom interface {
	int | string
}

func reverseGeneric2[T constraintCustom](slice []T) []T {

	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}