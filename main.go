package main

import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas["lais"] = 26
	pessoas["lucas"] = 31

	if idade, ok := pessoas["lucas"]; ok {
		fmt.Println("Pessoa existe", idade, ok)
	} else {
		fmt.Println("Pessoa nao existe")
	}
}
