package main

import (
	"fmt"
)

func main() {
	var (
		a    int
		num  int
		summ int
		flag bool
	)
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&num)
		if !flag {
			summ += num
			flag = !flag
		} else {
			summ -= num
			flag = !flag
		}
	}
	fmt.Println(summ)
}
