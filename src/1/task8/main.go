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
	fmt.Println((float64(length) + float64(height)) * (float64(length) + float64(height)))

}
