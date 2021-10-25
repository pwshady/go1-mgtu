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
	c, _ := strconv.Atoi(str1)

	if input.Scan() {
		str1 = input.Text()
	}

	rust, _ := strconv.Atoi(str1)
	mapper := make(map[string]int)

	for i := 0; i < c+rust; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		mapper[str1] = 1
	}

	fmt.Println(2*len(mapper) - c - rust)
}
