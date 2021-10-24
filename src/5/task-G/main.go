package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str1  string
		o_max int
		o_tec int
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}
	rune_str1 := []rune(str1)
	rune_str1 = append(rune_str1, 'p')

	for i := 0; i < len(rune_str1); i++ {
		if rune_str1[i] == 'о' || rune_str1[i] == 'О' {
			o_tec++
		} else {
			if o_tec >= o_max {
				o_max = o_tec
				o_tec = 0
			}
		}
	}
	fmt.Println(o_max)

}
