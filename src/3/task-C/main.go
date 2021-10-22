package main

import (
	"fmt"
)

func main() {
	var (
		a int64
	)

outer:
	for {
		fmt.Scan(&a)
		if a == 0 {
			break outer
		} else {
			fmt.Println(a)

		}
	}
}
