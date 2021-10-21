package main

import (
	"fmt"
)

func main() {
	var (
		x1 int
		x2 int
		y1 int
		y2 int
	)

	fmt.Scan(&x1, &y1, &x2, &y2)
	if (((x1-x2) == 2 || (x1-x2) == -2) && ((y1-y2) == 1 || (y1-y2) == -1)) || ((x1-x2) == 1 || (x1-x2) == -1 && ((y1-y2) == 2 || (y1-y2) == -2)) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}

}
