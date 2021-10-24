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
		summ int
	)

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str1 = input.Text()
	}

	count, _ := strconv.Atoi(str1)
	num_arr := [32]int{}

	for i := 0; i < count; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		num, _ := strconv.Atoi(str1)
		num_arr[i] = num
	}

	if input.Scan() {
		str1 = input.Text()
	}
	start, _ := strconv.Atoi(str1)
	if input.Scan() {
		str1 = input.Text()
	}
	end, _ := strconv.Atoi(str1)
	for i := start - 1; i < end; i++ {
		summ += num_arr[i]
	}
	fmt.Println(summ)
}
