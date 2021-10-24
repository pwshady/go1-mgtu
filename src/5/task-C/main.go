package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str1 string
		str2 string
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}

	for i := 0; i < 1000; i++ {
		if input.Scan() {
			str2 = input.Text()
		}
		rune_str1 := []rune(str1)
		rune_str2 := []rune(str2)
		if rune_str1[len(rune_str1)-1] != rune_str2[0] {
			fmt.Println(str2)
			return
		}
		str1 = str2
		i--
	}

}
