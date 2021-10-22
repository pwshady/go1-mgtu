package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		pass1 string
		pass2 string
	)

	for i := 0; i <= 5; {
		fmt.Scan(&pass1)
		fmt.Scan(&pass2)
		if len(pass1) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		}
		if strings.Contains(pass1, "123") || strings.Contains(pass1, "qwe") {
			fmt.Println("Слишком простой пароль!")
			continue
		}
		if pass1 != pass2 {
			fmt.Println("Введенные пароли различаются!")
			continue
		}
		fmt.Println("Ну наконец-то!")
		return
	}
}
