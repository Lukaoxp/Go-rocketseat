package main

func main() {
	println(soma(1, 2))

	multiplica := func(x int) int {
		return x * 2
	}

	resultado := multiplica(5)
	println(resultado)
}

func soma(a, b int) (retorno int) {
	retorno = a + b
	return
}
