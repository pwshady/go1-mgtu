package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		str1 string
		str2 string
		str3 string
	)

	fmt.Fscan(os.Stdin, &str1)
	fmt.Fscan(os.Stdin, &str2)
	fmt.Fscan(os.Stdin, &str3)
	if strings.Compare(str1, "раз") == 0 && strings.Compare(str2, "два") == 0 && strings.Compare(str3, "три") == 0 {
		fmt.Println("ОК")
		return
	}
	if strings.Compare(str1, "один") == 0 && strings.Compare(str2, "два") == 0 && strings.Compare(str3, "три") == 0 {
		fmt.Println("ОК")
		return
	}
	if strings.Compare(str1, "1") == 0 && strings.Compare(str2, "2") == 0 && strings.Compare(str3, "3") == 0 {
		fmt.Println("ОК")
		return
	}
	fmt.Println("НЕ ПРАВИЛЬНО")

}
