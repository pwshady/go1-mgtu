package main

import (
	"fmt"
)

func main() {
	var (
		puls  float64
		max   float64
		min   float64
		goden int64
	)

outer2:
	for {
		fmt.Scan(&puls)
		if puls >= 100 && puls <= 140 {
			goden++
			if max < puls {
				max = puls
			}
			if min > puls {
				min = puls
			}
			if min == 0 {
				min = puls
			}
		} else {
			if puls < 0 {
				break outer2
			}
		}
	}
	fmt.Println(goden)
	fmt.Printf("%.1f %.1f\n", min, max)
}
