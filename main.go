package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	CEP    string
}

func main() {
	cliente := Cliente{
		Nome:  "lucas",
		Idade: 31,
		Endereco: Endereco{
			CEP:    "88000-000",
			Numero: 100,
		},
	}

	cliente.Email = "teste@teste"
	fmt.Println(cliente)
}
