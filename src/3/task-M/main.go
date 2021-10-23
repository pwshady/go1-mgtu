package main

import (
	"fmt"
)

func main() {
	var (
		a    int
		x    float32
		y    float32
		xmin float32
		ymin float32
		xmax float32
		ymax float32
	)
	for i := 0; i < 4; i++ {
		fmt.Scan(&x, &y)
		if x > xmax {
			xmax = x
		}
		if y > ymax {
			ymax = y
		}
		if x < xmin {
			xmin = x
		}
		if y < ymin {
			ymin = y
		}
	}
	fmt.Scan(&a)
	arrx := [32]float32{}
	arry := [32]float32{}
	for i := 0; i < a; i++ {
		fmt.Scan(&arrx[i], &arry[i])
	}
	for i := 0; i < a; i++ {
		if arrx[i] >= xmin && arrx[i] <= xmax && arry[i] >= ymin && arry[i] <= ymax {
			fmt.Printf("Точка с координатами %.0f,%.0f принадлежит исследуемому квадрату\n", arrx[i], arry[i])
		} else {
			fmt.Printf("Точка с координатами %.0f,%.0f не принадлежит исследуемому квадрату\n", arrx[i], arry[i])
		}
	}

}
