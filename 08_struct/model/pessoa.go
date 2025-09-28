package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

// método
func (p *Pessoa) CalculaIdade() {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento
}

// função
func CalculaIdade(p Pessoa) int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
