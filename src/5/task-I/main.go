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

	count, _ := strconv.Atoi(str1)
	task_arr := [32]string{}

	for i := 0; i < count; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		task_arr[i] = str1
	}

	if input.Scan() {
		str1 = input.Text()
	}

	m, _ := strconv.Atoi(str1)
	m_arr := [32]int{}

	for i := 0; i < m; i++ {
		if input.Scan() {
			str1 = input.Text()
		}
		count, _ := strconv.Atoi(str1)
		m_arr[i] = count
	}

	for i := 0; i < 32; i++ {
		if m_arr[i] == 0 {
			return
		}
		fmt.Println(task_arr[m_arr[i]-1])
	}
	fmt.Println(len(m_arr))

}
