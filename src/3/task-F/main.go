package main

import (
	"fmt"
)

func main() {
	var (
		number  int64
		num_str string
	)
	num_str = ""
	fmt.Scan(&number)
	for number >= 1 {
		if number%2 == 1 {
			num_str = "1" + num_str
			number--
		} else {
			num_str = "0" + num_str
		}
		number = number / 2
	}
	fmt.Println(num_str)
}
