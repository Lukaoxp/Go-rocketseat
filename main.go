package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Apresentar() {
	fmt.Printf("Olá, sou %s e tenho %d anos.\n", p.Nome, p.Idade)
	p.Nome = "mario"
}

func main() {
	pessoa := Pessoa{
		Nome:  "Lucas",
		Idade: 31,
	}

	pessoa.Apresentar()
	pessoa.Apresentar()
}
