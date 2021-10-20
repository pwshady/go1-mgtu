package main

import (
	"fmt"
)

func main() {
	var (
		num1 string
		num2 string
		num3 string
	)

	fmt.Scan(&num1, &num2, &num3)
	fmt.Println(num3 + num2 + num1)

}
