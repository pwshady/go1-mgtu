package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str1 string
		flag bool
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}
	rune_str1 := []rune(str1)

	for i := 0; i < len(rune_str1); i++ {
		rune_str2 := []rune("")
		fmt.Println(string(rune_str1))
		if !flag {
			for j := 2; j < len(rune_str1); j++ {
				rune_str2 = append(rune_str2, rune_str1[j])
			}
		} else {
			for j := 0; j < len(rune_str1)-2; j++ {
				rune_str2 = append(rune_str2, rune_str1[j])
			}
		}
		rune_str1 = rune_str2
		flag = !flag
		i--
	}

}
