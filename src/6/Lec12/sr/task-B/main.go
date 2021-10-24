package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		str1 string
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}

	mass, _ := strconv.Atoi(str1)

	for i := 0; i < mass; i++ {
		for j := i; j < mass-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("@")
		}
		fmt.Print("@")
		for j := 0; j < i; j++ {
			fmt.Print("@")
		}
		fmt.Println("")
	}

}
