package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var (
		logo  string
		price int
		rub   int
		kop   int
	)

	fmt.Fscan(os.Stdin, &logo)
	price = utf8.RuneCountInString(logo) * 23
	kop = price % 100
	rub = int(price / 100)
	if rub == 0 {
		fmt.Println(kop, "коп.")
		return
	}
	fmt.Println(rub, "р.", kop, "коп.")

}
