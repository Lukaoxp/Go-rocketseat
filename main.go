package main

func main() {

	// for normal
	for i := 0; i < 10; i++ {
		// println(i)
	}

	// ou
	for i := range 10 {
		println(i)
	}

	println()

	// for como um while
	sum := 1
	for sum < 20 {
		sum += sum
		println(sum)
	}

	println()
	// for com length de um Slice
	nums := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	// //for infinito
	// for {
	// 	println("oi")
	// }

	println()
	// com range, consigo ter a key e value
	for key, value := range nums {
		println(key, value)
	}
}
