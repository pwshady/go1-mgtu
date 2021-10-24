package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	count, _ := strconv.Atoi(str1)
	mapper := make(map[string]int)

	for i := 0; i < count; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		mapper[str1] = 0
	}

	if input.Scan() {
		str1 = input.Text()
	}

	count, _ = strconv.Atoi(str1)

	for i := 0; i < count; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		count_tec, _ := strconv.Atoi(str1)
		for j := 0; j < count_tec; j++ {
			if input.Scan() {
				str1 = input.Text()
			}
			mapper[str1] += 1
		}
	}

	result_mapper := make(map[string]int)
	for key, value := range mapper {
		if value == 0 {
			result_mapper[key] = 1
		}
	}
	//Sort
	keys := make([]string, 0, len(result_mapper))
	for k := range result_mapper {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//
	for _, k := range keys {
		fmt.Println(k)
	}
}
