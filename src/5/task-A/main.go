package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		str  string
		flag bool
	)

	arr_all := [32]string{}
	my_arr := [32]string{}

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		str = input.Text()
	}
	all, _ := strconv.Atoi(str)

	if input.Scan() {
		str = input.Text()
	}
	my_film, _ := strconv.Atoi(str)

	for i := 0; i < all; i++ {
		if input.Scan() {
			arr_all[i] = input.Text()
		}
	}

	for i := 0; i < my_film; i++ {
		if input.Scan() {
			my_arr[i] = input.Text()
		}
	}

	for i := 0; i < my_film; i++ {
		for j := 0; j < all; j++ {
			if my_arr[i] == arr_all[j] {
				fmt.Println("ЕСТЬ")
				flag = true
			}
		}
		if !flag {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}

}
