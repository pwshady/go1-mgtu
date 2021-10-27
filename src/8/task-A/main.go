package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fact(a int) int {
	var b int = 1
	for i := 1; i <= a; i++ {
		b = b * i
	}
	return b
}

func comb(n int, m int) int {
	return int(fact(n) / (fact(m) * fact(n-m)))
}

func GetIntParam() (int, error) {
	input := bufio.NewScanner(os.Stdin)
	var str string
	if input.Scan() {
		str = input.Text()
	}
	return strconv.Atoi(str)
}

func main() {
	n, _ := GetIntParam()
	m, _ := GetIntParam()
	fmt.Println(comb(n, m))
}
