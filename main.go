package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("hoje")
	case today + 1:
		fmt.Println("amanha")
	case today + 2:
		fmt.Println("depois de amanha")
	default:
		fmt.Println("Ta longe")
	}
}
