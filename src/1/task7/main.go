package main

import (
	"fmt"
)

func main() {
	var (
		length float32
		height float32
	)

	fmt.Scan(&length)
	fmt.Scan(&height)
	fmt.Println("Периметр:", (2 * (height + length)))
	fmt.Println("Площадь:", height*length)

}
