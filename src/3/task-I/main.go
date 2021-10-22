package main

import (
	"fmt"
)

func main() {
	var (
		num    int
		ahtung bool
	)
	fmt.Scan(&num)
	ahtung = true
	if num == 1 {
		fmt.Println("1")
		return
	}
	fmt.Print("1")
outer:
	for i := 2; i <= int(num); i++ {
		if num%i == 0 {
			if i != num {
				ahtung = false
			}
			fmt.Print(" ", i)
		}
		if i >= 1000000 {
			break outer
		}
	}
	if ahtung {
		fmt.Println("")
		fmt.Println("ACHTUNG")
	}
}
