package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		a string
	)

	var loopVar int = 0
	for loopVar < 10 {
		fmt.Scanln(&a)
		if len(a) == 0 {
			fmt.Scanln(&a)
			break
		}

		if strings.Contains(a, "") {
			fmt.Println(a)
			a = ""
		}
	}
}
