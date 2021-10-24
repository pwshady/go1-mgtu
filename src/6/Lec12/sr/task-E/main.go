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
		two  int
		flag bool
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

	c_arr := []string{}
	rust_arr := []string{}

	for i := 0; i < c; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		c_arr = append(c_arr, str1)
	}

	for i := 0; i < rust; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		rust_arr = append(rust_arr, str1)
	}

	for i := 0; i < c; i++ {
		for j := 0; j < rust; j++ {
			if c_arr[i] == rust_arr[j] {
				two++
			}
		}

	}

	for i := 0; i < c; i++ {
		flag = false
		for j := 0; j < rust; j++ {
			if c_arr[i] == rust_arr[j] {
				flag = true
			}
		}
		if !flag {
			rust_arr = append(rust_arr, c_arr[i])
		}
	}

	fmt.Println(len(rust_arr) - two)

}
