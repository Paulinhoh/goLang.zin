package main

import (
	"errors"
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())
}

func main() {
	fmt.Println("Iniciando...")

	retangulo := retangulo{
		largura: 1,
		altura:  2,
	}
	circulo := circulo{
		radius: 3,
	}

	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)

	// =-=-=-=-=-=-=-=-=-=-=-=-= Outro exemplo =-=-=-=-=-=-=-=-=-=-=-=-=
	p := ProblemaDeNetwork{
		rede:     true,
		hardware: false,
	}

	ExibirError(errors.New("a error"))
	ExibirError(p)

	// =-=-=-=-=-=-=-=-=-=-=-=-= Interface vazia == interface{} =-=-=-=-=-=-=-=-=-=-=-=-=
	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, true)
	lista = append(lista, "teste")

	for _, valor := range lista {

		if v, ok := valor.(string); ok {
			fmt.Println(v + " (string)")
		} else {
			fmt.Println(valor)
		}
	}
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "Problema de Rede"
	} else if p.hardware {
		return "Problema de Hardware"
	} else {
		return "Outro Problema"
	}
}

func ExibirError(err error) {
	fmt.Println(err.Error())
}
