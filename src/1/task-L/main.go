package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		log  string
		pass string
	)

	fmt.Fscan(os.Stdin, &log)
	fmt.Fscan(os.Stdin, &pass)
	//fmt.Fscan(os.Stdin, &subStr)
	if strings.Contains(log, "@") || utf8.RuneCountInString(log) < 10 {
		fmt.Println("Некорректный логин")
		return
	} else {
		if !strings.Contains(pass, "@") || !strings.Contains(pass, ".") {
			fmt.Println("Некорректная почта")
			return
		}
	}
	fmt.Println("ОК")

}
