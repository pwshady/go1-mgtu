package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str string
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str = input.Text()
	}

	rune_str := []rune(str)
	str = string(rune_str[0]) + string(rune_str[len(rune_str)-1])

	for i := 0; i < 2; i++ {
		if str == "Да" {
			fmt.Println("СОГЛАСЕН")
			return
		}
		if str == "дА" {
			fmt.Println("СОГЛАСЕН")
			return
		}
		if str == "ДА" {
			fmt.Println("СОГЛАСЕН")
			return
		}
		if str == "да" {
			fmt.Println("СОГЛАСЕН")
			return
		}
		str = string(rune_str[len(rune_str)-1]) + string(rune_str[0])
	}
	fmt.Println("НЕ СОГЛАСЕН")

}
