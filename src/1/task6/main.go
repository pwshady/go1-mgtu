package main

import (
	"fmt"
)

func main() {
	var (
		mes1 string
		mes2 string
		mes3 string
		mes4 string
		name string
	)

	fmt.Scan(&mes1, &mes2, &mes3, &mes4)
	fmt.Scan(&name)
	fmt.Println(mes4, "-", name)
	fmt.Println(mes3, "-", name)
	fmt.Println(mes2, "-", name)
	fmt.Println(mes1, "-", name)
}
