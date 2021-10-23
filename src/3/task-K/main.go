package main

import (
	"fmt"
)

func main() {
	var (
		num int
	)
	fmt.Scan(&num)
	for i := 0; i < num-1; i++ {
		fmt.Print(" ")
	}
	if num >= 1 {
		fmt.Println("1")
	}
	for i := 0; i < num-2; i++ {
		fmt.Print(" ")
	}
	if num >= 2 {
		fmt.Println("1 1")
	}
	arr1 := [32]int{1}
	arr2 := [32]int{1, 1}
outer:
	for i := 2; i < int(num); i++ {
		for j := 0; j < len(arr1)-1; j++ {
			arr1[j+1] = arr2[j] + arr2[j+1]
		}
		arr1[i] = 1
		arr2 = arr1
		for k := 0; k < num-i-1; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < 32; j++ {
			if arr1[j] != 0 {
				if j+1 == num {
					fmt.Print(arr1[j])
				} else {
					fmt.Print(arr1[j], " ")
				}

			}
		}
		fmt.Println()
		if i >= 32 {
			break outer
		}
	}
}
