package main

import (
	"fmt"
)

func main() {
	var (
		num int
		st1 array
	)
	fmt.Scan(&num)
	if num < 0 {
		num = 0 - num

	}
	int_num = int(num)
outer:
	for i := -int_num; i <= int(num); i++ {
		fmt.Println("Квадрат числа", i, "равен", i*i)
		if i >= 1000000 {
			break outer
		}
	}
}
