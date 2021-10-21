package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		str    string
		subStr string = "чат"
	)

	fmt.Fscan(os.Stdin, &str)
	//fmt.Fscan(os.Stdin, &subStr)
	if strings.Contains(str, subStr) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}

}
