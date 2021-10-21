package main

import (
	"fmt"
)

func main() {
	var (
		a float32
		b float32
		c float32
	)

	fmt.Scan(&a, &b, &c)
	if c == 0 {
		fmt.Println("один корень")
		return
	}

	if b == 0 && c < 0 {
		fmt.Println("два корня")
		return
	}

	if b == 0 && c > 0 {
		fmt.Println("корней нет")
		return
	}

	if a == 0 || b*b-4*a*c == 0 {
		fmt.Println("один корень")
		return
	}

	if b*b-4*a*c < 0 {
		fmt.Println("корней нет")
		return
	}
	fmt.Println("два корня")

}
