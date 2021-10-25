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
	if mass > 3 {
		if mass%2 == 0 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")

}
