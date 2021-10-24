package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str1 string
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}
	rune_str1 := []rune(str1)
	rune_str2 := []rune("")

	for i := 0; i < len(rune_str1); i += 2 {
		for j := 0; j < 3; j++ {
			rune_str2 = append(rune_str2, rune_str1[i])
		}
	}

	fmt.Println(string(rune_str2))

}
