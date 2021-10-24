package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		kr   int
		str1 string
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}
	summ, _ := strconv.Atoi(str1)

	if input.Scan() {
		str1 = input.Text()
	}
	st, _ := strconv.Atoi(str1)

	kr = int(summ / 5)
	summ = kr * 5
	if summ >= 20*(st-2) {
		fmt.Println(st, 0, 0)
		return
	}
	for i := 1; i <= kr; i++ {
		for j := 0; j < kr-i; j++ {
			for k := 0; k < kr-i-j; k++ {
				if (i*20+j*10+k*5) == summ && i+j+k == st {
					fmt.Println(i, j, k)
				}
			}

		}
	}
}
