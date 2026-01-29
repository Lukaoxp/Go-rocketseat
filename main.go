package main

import "fmt"

func main() {
	var nota int

	fmt.Print("Digite a nota: ")
	fmt.Scanln(&nota)

	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	} else if nota >= 70 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
