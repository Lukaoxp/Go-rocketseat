package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := retornarErro(); err != nil {
		fmt.Println(err.Error())
	}

	players := map[string]int{
		"lucas": 31,
		"mario": 14,
	}

	if value, ok := players["lucas"]; ok {
		fmt.Println("pontos: ", value, ok)
	}

}

func retornarErro() error {
	return errors.New("Isso é um erro")
}
