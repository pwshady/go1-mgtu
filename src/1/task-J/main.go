package main

import (
	"fmt"
)

func main() {
	var (
		num1 float32
		num2 float32
		summ float32
	)

	fmt.Scan(&num1, &num2)
	summ = num1 + num2
	if int(summ)%2 == 1 {
		fmt.Println("НЕЧЁТНОЕ")
	} else {
		fmt.Println("ЧЁТНОЕ")
	}

}
